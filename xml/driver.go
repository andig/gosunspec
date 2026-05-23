package xml

import (
	"errors"
	"github.com/andig/gosunspec"
	"github.com/andig/gosunspec/impl"
	"github.com/andig/gosunspec/models/model1"
	"github.com/andig/gosunspec/types"
	"github.com/andig/gosunspec/spi"
	"github.com/andig/gosunspec/typelabel"
	_ "log"
	"strconv"
	"strings"
)

type blockAnchor struct {
	device   *DeviceElement
	model    *ModelElement
	index    int
	detached bool
}

type pointAnchor struct {
	position int // position of the point in the model, -1 if it is not present
}

type xmlDriver struct {
}

var ErrNoSuchElement = errors.New("no such element")
var ErrEmptyValue = errors.New("empty value")
var ErrTooManyErrors = errors.New("too many errors")

// Open open a new sunspec.Array which is populated from the specified
// DataElement.
func Open(e *DataElement) (sunspec.Array, error) {
	arr := impl.NewArray()
	for x, _ := range e.Devices {
		if dev, err := OpenDevice(e.Devices[x]); err != nil {
			return nil, err
		} else if err := arr.AddDevice(dev.(spi.DeviceSPI)); err != nil {
			return nil, err
		}
	}
	return arr, nil
}

// OpenDevice opens a new sunspec.Device which is populated from the specified
// DeviceElement.
func OpenDevice(dx *DeviceElement) (sunspec.Device, error) {
	d := impl.NewDevice()
	d.SetAnchor(dx)
	xp := &xmlDriver{}

	//
	// Special case for optional model 1 element in XML representation
	//
	// The SunSpec Model Data Exchange specification allows the model 1
	// element to be omitted from device documents in certain circumstances.
	//
	// In these cases several of the common model elements might be
	// exchanged instead as attributes of the device element. To allow
	// a generic API to bs used to read and write these points, it
	// is simplest to create a detached ModelElement for this case.
	//
	// We don't link the model element to the parent device element unless
	// an attempt is made to write a point that is not duplicated by
	// a device element attribute.
	//
	// The intent is to provided a consistent API across XML and non-XML
	// drivers and to prevent model 1 elements magically appearing in the
	// XML device element merely because an XML device element has been
	// opened but also to ensure that if a write is made into the model 1
	// element, that write is preserved and visible to the consumers
	// of the XML model.
	//
	foundModel1 := false
	for _, mx := range dx.Models {
		foundModel1 = (sunspec.ModelId(mx.Id) == model1.ModelID)
		if foundModel1 {
			break
		}
	}

	models := dx.Models
	if !foundModel1 {
		models = make([]*ModelElement, len(dx.Models)+1)
		models[0] = &ModelElement{Id: model1.ModelID}
		copy(models[1:], dx.Models)
	}

	// iterate through the model elements, creating one new model for each
	// then create a block for each index and add points for each point element
	for _, mx := range models {
		md := types.GetModel(uint16(mx.Id))
		if md == nil {
			continue
		}
		max := uint32(0)
		for _, px := range mx.Points {
			if px.Index > max {
				max = px.Index
			}
		}
		repeats := int(max)
		if len(md.Blocks) == 1 {
			repeats -= 1
		}
		m := impl.NewModel(md, repeats, xp)
		if err := d.AddModel(m); err != nil {
			return nil, err
		} else {
			for pi, px := range mx.Points {
				bi := int(px.Index)
				if bi > 0 && len(md.Blocks) == 1 {
					bi = bi - 1
				}
				b := m.MustBlock(bi)
				if p, err := b.Point(px.Id); err != nil {
					return nil, err
				} else {
					p.(spi.PointSPI).SetAnchor(&pointAnchor{position: pi})
				}
			}
			bi := 0
			m.Do(spi.WithBlockSPI(func(b spi.BlockSPI) {
				detached := mx.Id == model1.ModelID && !foundModel1
				b.SetAnchor(&blockAnchor{device: dx, model: mx, index: bi, detached: detached})
				b.Do(spi.WithPointSPI(func(p spi.PointSPI) {
					if p.Anchor() == nil {
						p.SetAnchor(&pointAnchor{position: -1})
					}
				}))
				bi++
			}))
		}
	}

	if !foundModel1 {
		syncDeviceElementToModel1(dx, d)
	}

	return d, nil
}

//
// For a fake model1 element, initialise it to be consistent with the device element
//
func syncDeviceElementToModel1(dx *DeviceElement, d sunspec.Device) {
	b := d.MustModel(model1.ModelID).MustBlock(0)
	b.MustPoint(model1.SN).SetStringValue(dx.Serial)
	b.MustPoint(model1.Md).SetStringValue(dx.Model)
	b.MustPoint(model1.Mn).SetStringValue(dx.Manufacturer)
	if err := b.Write(model1.SN, model1.Md, model1.Mn); err != nil {
		// we are not expecting this
		panic(err)
	}
}

// CopyArray copies an existing SunSpec Array into a new SunSpec Array and an
// XML DataElement. Operations on the returned SunSpec Array edit the
// returned DataElement.
func CopyArray(a sunspec.Array) (sunspec.Array, *DataElement) {
	arr := impl.NewArray()
	x := &DataElement{
		Version: "1",
	}

	a.Do(func(d sunspec.Device) {
		da, dx := CopyDevice(d)
		x.Devices = append(x.Devices, dx)
		if err := arr.AddDevice(da.(spi.DeviceSPI)); err != nil {
			// we are not expecting this to happen
			panic(err)
		}
	})
	return arr, x
}

// CopyDevice copies an existing SunSpec Device into a new SunSpec Device and an
// XML DeviceElement. Operations on the returned SunSpec Device edit the
// returned DeviceElement.
func CopyDevice(d sunspec.Device) (sunspec.Device, *DeviceElement) {
	dc := impl.NewDevice()
	dx := &DeviceElement{}
	xp := &xmlDriver{}

	modelElements := map[sunspec.ModelId]*ModelElement{}

	newModelElement := func(id sunspec.ModelId) *ModelElement {

		// we only use labels when there is ambiguity

		var mx *ModelElement
		var ok bool
		if mx, ok = modelElements[id]; ok {
			if mx.Index == 0 {
				mx.Index = 1
			}
			mx = &ModelElement{Id: id, Index: mx.Index + 1}
		} else {
			mx = &ModelElement{Id: id, Index: 0}
		}
		modelElements[id] = mx
		return mx
	}

	d.Do(func(m sunspec.Model) {
		md := types.GetModel(uint16(m.Id()))

		repeatOnly := m.Blocks() > 1 && len(md.Blocks) == 1
		mc := impl.NewModel(md, m.Blocks()-1, xp)
		mx := newModelElement(m.Id())

		dx.Models = append(dx.Models, mx)

		for i := 0; i < m.Blocks(); i++ {
			b := m.MustBlock(i).(spi.BlockSPI)
			bc := mc.MustBlock(i).(spi.BlockSPI)

			x := i
			if repeatOnly {
				x = x + 1
			}
			anchor := &blockAnchor{
				device: dx,
				model:  mx,
				index:  x,
			}
			bc.SetAnchor(anchor)
			points, _ := bc.Plan()
			for _, pc := range points {
				pa := &pointAnchor{position: -1}
				pc.SetAnchor(pa)
				p := b.MustPoint(pc.Id())
				if p.Error() == nil {
					pc.SetValue(p.Value())
				}
			}

			// write the copied values into the copied block

			if err := bc.Write(); err != nil {
				// not expecting this to happen
				panic(err)
			}
		}
		dc.AddModel(mc)
	})
	return dc, dx
}

func deriveScaleFactor(v string, sf int16) (string, sunspec.ScaleFactor) {
	pos := strings.Index(v, ".")
	if pos < 0 {
		return v, sunspec.ScaleFactor(sf)
	} else {
		// "1.3", pos=1, len=3 -> 1-3+1 -> 2-3 -> -1
		return strings.Replace(v, ".", "", 1), sunspec.ScaleFactor(pos - len(v) + 1)
	}
}

func (phys *xmlDriver) Read(b spi.BlockSPI, pointIds ...string) error {
	errCount := 0
	var firstError error
	ba := b.Anchor().(*blockAnchor)
	if points, err := b.Plan(pointIds...); err != nil {
		return err
	} else {
		for _, p := range points {
			recordError := func(e error) {
				if firstError == nil || firstError == ErrNoSuchElement {
					firstError = e
				}
				p.SetError(e)
				errCount++
				// log.Printf("point error: id=%s: %v", p.Id(), e)
			}

			pa := p.Anchor().(*pointAnchor)
			if pa.position < 0 {
				recordError(ErrNoSuchElement)
				continue
			}
			px := ba.model.Points[pa.position]
			if len(px.Value) == 0 && p.Type() != typelabel.String {
				recordError(ErrEmptyValue)
				continue
			}
			sfp := p.ScaleFactorPoint()
			v := px.Value
			if sfp != nil {
				var vsf sunspec.ScaleFactor
				v, vsf = deriveScaleFactor(v, px.ScaleFactor)
				if sfp.Error() != nil {
					sfp.SetScaleFactor(vsf)
				}
				if vi, err := strconv.Atoi(v); err != nil {
					recordError(err)
					continue
				} else {
					for vsf > sfp.ScaleFactor() {
						vi = vi * 10
						vsf--
					}
					for vsf < sfp.ScaleFactor() {
						vi = vi / 10
						vsf++
					}
					v = strconv.Itoa(vi)
				}
			}
			if err := p.UnmarshalXML(v); err != nil {
				recordError(err)
				continue
			}

			// need to values with scaling factors match the scaling factor or else
			// adjust the value to be consistent with the current scaling factor value
		}
	}
	if errCount > 0 {
		return firstError
	} else {
		return nil
	}
}

func (phys *xmlDriver) Write(b spi.BlockSPI, pointIds ...string) error {

	write := map[string]bool{}
	if len(pointIds) == 0 {
		b.Do(func(p sunspec.Point) {
			write[p.Id()] = true
		})
	} else {
		for _, id := range pointIds {
			write[id] = true
		}
	}

	ba := b.Anchor().(*blockAnchor)
	b.Do(spi.WithPointSPI(func(p spi.PointSPI) {
		if !write[p.Id()] {
			return
		}
		pa := p.Anchor().(*pointAnchor)
		var px *PointElement
		if pa.position == -1 && p.Error() == nil {
			// add a new point to the model element
			px = &PointElement{Id: p.Id(), Index: uint32(ba.index)}
			pa.position = len(ba.model.Points)
			ba.model.Points = append(ba.model.Points, px)
		} else if pa.position >= 0 && p.Error() != nil {
			// remove a point from the model element
			swap := b.MustPoint(ba.model.Points[len(ba.model.Points)-1].Id).(spi.PointSPI).Anchor().(*pointAnchor)
			ba.model.Points[pa.position] = ba.model.Points[swap.position]
			swap.position = pa.position
			ba.model.Points = ba.model.Points[0 : len(ba.model.Points)-1]
			px = nil
		} else if pa.position >= 0 {
			// update the existing point
			px = ba.model.Points[pa.position]
		}
		if px != nil {
			px.Value = p.MarshalXML()
			if sfp := p.ScaleFactorPoint(); sfp != nil {
				px.ScaleFactor = int16(sfp.ScaleFactor())
			}
			// we could put the description and units (from the SMDX model) in here, but
			// we probably need a configuration option to say whether this is necessary
			// or not
		}
	}))

	// A special case required by a partial redundancy of the XML model.
	//
	// Writes into the model 1 points that also appear in the device element
	// as attributes are mirrored into the device element attributes.
	//
	// In all other cases, we need to link a detached model element to the
	// parent so that these updates are visible to consumers of the XML model.
	//
	// See also: OpenDevice

	if ba.model.Id == model1.ModelID {

		count := 0
		for pid, _ := range write {
			switch pid {
			case model1.SN, model1.Md, model1.Mn:
				// synchronize device element with writes into model
				p := b.MustPoint(pid)
				var v string
				if p.Error() == nil {
					v = p.StringValue()
				} else {
					v = ""
				}
				switch pid {
				case model1.SN:
					ba.device.Serial = v
				case model1.Md:
					ba.device.Model = v
				case model1.Mn:
					ba.device.Manufacturer = v
				}
			default:
				// count writes into other points
				count++
			}
		}

		// If there is a write to a point that isn't encoded in the device
		// element and the model 1 element is currently detached from the device
		// (because it wasn't present on open) then fix that by attaching the
		// model element to the parent device now.

		if count > 0 && ba.detached {
			models := make([]*ModelElement, len(ba.device.Models)+1)
			models[0] = ba.model
			copy(models[1:], ba.device.Models)
			ba.detached = false
			ba.device.Models = models
		}
	}

	return nil
}
