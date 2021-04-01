package layout

import (
	"encoding/binary"
	"log"
	"net"

	"github.com/andig/gosunspec/impl"
	"github.com/andig/gosunspec/models/model1"
	"github.com/andig/gosunspec/smdx"
	"github.com/andig/gosunspec/spi"
)

// SunspecLayout is the type of layout that understands the SunSpec layout conventions.
type SunSpecLayout struct {
}

// Open scans the supplied address space and returns an array of the
// devices found.
func (s *SunSpecLayout) Open(a AddressSpaceDriver) (spi.ArraySPI, error) {

	// Attempt to locate SunSpec register within modbus address space.

	baseRange := a.BaseOffsets()
	base := uint16(0xffff)
	for _, b := range baseRange {
		if id, err := a.ReadWords(b, 2); err != nil {
			// if one query fails with network error including timeout,
			// then assume they all will.
			if err, ok := err.(net.Error); ok {
				return nil, err
			}
			continue
		} else if binary.BigEndian.Uint32(id) != SunSpec {
			continue
		} else {
			base = b
			break
		}
	}
	if base == 0xffff {
		return nil, ErrNotSunspecDevice
	}

	phys := a
	array := impl.NewArray()
	dev := impl.NewDevice()

	// Build up model

	offset := uint16(2) // number of 16 bit registers
	for {
		if bytes, err := a.ReadWords(base+offset, 2); err != nil {
			if offset > 2 { // model chain partially available
				return array, err
			}
			return nil, err
		} else if len(bytes) < 4 {
			return nil, ErrShortRead
		} else {
			modelId := binary.BigEndian.Uint16(bytes)
			modelLength := binary.BigEndian.Uint16(bytes[2:])

			if modelId == 0xffff {
				break
			}

			me := smdx.GetModel(modelId)
			if me != nil {

				if modelId == uint16(model1.ModelID) {
					dev = impl.NewDevice()
					array.(spi.ArraySPI).AddDevice(dev)
				}

				m := impl.NewContiguousModel(me, modelLength, phys)

				// set anchors on the blocks

				blockOffset := offset + 2
				m.Do(spi.WithBlockSPI(func(b spi.BlockSPI) {
					b.SetAnchor(uint16(base + blockOffset))
					blockOffset += b.Length()
				}))
				dev.AddModel(m)
			} else {
				log.Printf("unrecognised model identifier skipped @ offset: %d, %d\n", modelId, offset)
			}
			offset += 2 + modelLength
		}
	}
	return array, nil

}
