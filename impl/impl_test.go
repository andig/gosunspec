package impl

import (
	"bytes"
	"math"
	"testing"

	sunspec "github.com/andig/gosunspec"
	"github.com/andig/gosunspec/smdx"
	"github.com/andig/gosunspec/spi"
	"github.com/andig/gosunspec/typelabel"
)

func TestCompletePointInterface(t *testing.T) {
	_ = spi.PointSPI((&point{}))
}

func TestCompleteDevice(t *testing.T) {
	_ = spi.DeviceSPI((&device{}))
}

func TestCompleteArray(t *testing.T) {
	_ = spi.ArraySPI((&array{}))
}

func TestNotImplemented(t *testing.T) {
	cases := []struct {
		expect bool
		value  interface{}
	}{
		{false, float32(0)},
		{true, math.Float32frombits(0x7fc00000)},
		{false, int16(0)},
		{false, int32(0)},
		{false, int64(0)},
		{true, int16(-0x8000)},
		{true, int32(-0x80000000)},
		{true, int64(-0x8000000000000000)},
		{false, uint16(0)},
		{false, uint32(0)},
		{false, uint64(0)},
		{true, uint16(0xFFFF)},
		{true, uint32(0xFFFFFFFF)},
		{true, uint64(0xFFFFFFFFFFFFFFFF)},
		{true, sunspec.Acc16(0xFFFF)},
		{true, sunspec.Acc32(0xFFFFFFFF)},
		{true, sunspec.Acc64(0xFFFFFFFFFFFFFFFF)},
		{false, sunspec.Ipaddr{0xFF, 0xFF, 0xFF, 0xFF}},
		{false, sunspec.Ipv6addr{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}},
		{true, sunspec.Ipaddr{0, 0, 0, 0}},
		{true, sunspec.Ipv6addr{0, 0, 0, 0, 0, 0}},
		{false, sunspec.Eui48{0, 0, 0, 0, 0, 0}},
		{true, sunspec.Eui48{0x01, 0x02, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}},
		{false, sunspec.Bitfield16(0)},
		{false, sunspec.Bitfield32(0)},
		{true, sunspec.Bitfield16(0xFFFF)},
		{true, sunspec.Bitfield32(0xFFFFFFFF)},
		{false, sunspec.Enum16(0)},
		{false, sunspec.Enum32(0)},
		{true, sunspec.Enum16(0xFFFF)},
		{true, sunspec.Enum32(0xFFFFFFFF)},
	}

	for _, c := range cases {
		p := point{
			err:   nil,
			value: c.value,
		}

		if exp := p.NotImplemented(); c.expect != exp {
			t.Errorf("expected %v, got %v for %v", c.expect, exp, c.value)
		}

		if v, ok := p.Value().(sunspec.NotImplemented); !ok {
			t.Errorf("expected sunspec.NotImplemented, got %v", v)
		}
	}
}

func TestMarshalEui48(t *testing.T) {
	p := point{
		err: nil,
		smdx: &smdx.PointElement{
			Type: typelabel.Eui48,
		},
	}

	p.Unmarshal([]byte{1, 2, 3, 4, 5, 6, 7, 8})

	v := p.Value().(sunspec.Eui48)
	if !bytes.Equal(v[:], []byte{1, 2, 3, 4, 5, 6, 7, 8}) {
		t.Errorf("unexpected value result, got %v", v)
	}

	if v := p.MarshalXML(); v != "03:04:05:06:07:08" {
		t.Errorf("unexpected xml result, got %v", v)
	}
}
