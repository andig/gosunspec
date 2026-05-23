package layout

import (
	"encoding/xml"
	"errors"
	"github.com/andig/gosunspec"
	"github.com/andig/gosunspec/impl"
	"github.com/andig/gosunspec/types"
	"github.com/andig/gosunspec/spi"
	"io"
	"log"
)

var ErrAbsoluteAddress = errors.New("The absolute address for block cannot be calculated from the layout.")

// RawModel layout specifies the model ID used to describe a block of memory at an Address
type RawModelLayout struct {
	XMLName xml.Name        `xml:"model"`
	ModelId sunspec.ModelId `xml:"id,attr"`
	Address *uint16         `xml:"addr,attr,omitempty"`
	Repeats *uint16         `xml:"repeats,attr,omitempty"`
}

// RawDeviceLayout is a slice of RawModelLayouts
type RawDeviceLayout struct {
	XMLName xml.Name         `xml:"device"`
	Models  []RawModelLayout `xml:"model"`
}

// RawLayout is the type of layout used for non-SunSpec address spaces. This means
// arbitrary Modbus address spaces where blocks are located at arbitrary addresses
// in an address space and neither the model ID nor the block length are encoded
// in the address space itself.
//
// The intent of RawLayout is to allow the sunspec API to be used effectively
// with non-SunSpec address spaces, assuming the work has been done to
// markup the address space with SMDX models and an XML document that maps
// addresses to model ids.
type RawLayout struct {
	XMLName xml.Name          `xml:"layout"`
	Name    string            `xml:"name,attr,omitempty"`
	Devices []RawDeviceLayout `xml:"device"`
}

// FromLayoutXML reads a layout description from an XML stream.
func FromLayoutXML(reader io.Reader) (*RawLayout, error) {
	decoder := xml.NewDecoder(reader)
	layout := &RawLayout{}
	return layout, decoder.Decode(layout)
}

// Opens an address space with a raw layout.
func (s *RawLayout) Open(driver AddressSpaceDriver) (spi.ArraySPI, error) {

	array := impl.NewArray()
	dev := impl.NewDevice()

	// Build up model

	var nextAddr *uint16

	for _, d := range s.Devices {
		for _, m := range d.Models {

			me := types.GetModel(uint16(m.ModelId))
			if me != nil {

				modelLength := me.Blocks[0].Length
				r := uint16(0)
				if m.Repeats != nil && *m.Repeats > 0 {
					r = *m.Repeats
					if len(me.Blocks) > 1 {
						modelLength += me.Blocks[1].Length * r
					} else {
						modelLength += me.Blocks[0].Length * r
					}
				}

				model := impl.NewContiguousModel(me, modelLength, driver)

				// set anchors on the blocks
				var offset uint16

				if m.Address != nil {
					offset = *m.Address
				} else if nextAddr != nil {
					offset = *nextAddr
				} else {
					return nil, ErrAbsoluteAddress
				}

				model.Do(spi.WithBlockSPI(func(b spi.BlockSPI) {
					b.SetAnchor(uint16(offset))
					offset += b.Length()
				}))

				nextAddr = &offset
				dev.AddModel(model)
			} else {
				nextAddr = nil
				log.Printf("unrecognised model identifier skipped: %d\n", m.ModelId)
			}
		}
		array.AddDevice(dev)
	}
	return array, nil
}
