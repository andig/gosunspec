package impl

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"math"
	"strconv"

	sunspec "github.com/andig/gosunspec"
	"github.com/andig/gosunspec/smdx"
	"github.com/andig/gosunspec/spi"
	"github.com/andig/gosunspec/typelabel"
	"github.com/andig/gosunspec/typelen"
)

var (
	errNotInitialized = errors.New("point not initialised")
	errBadType        = errors.New("bad type")
	errLength         = errors.New("length")
	errValueTooLarge  = errors.New("value too large")
)

type point struct {
	anchored
	smdx        *smdx.PointElement
	scaleFactor sunspec.Point
	block       *block
	err         error
	value       interface{}
}

// Checks that the point does not have an error
func (p *point) checkerror() interface{} {
	if err := p.Error(); err != nil {
		panic(err)
	}
	return p.value
}

// Checks that the receiver has the specified typed
func (p *point) checktype(t string, v interface{}) {
	if p.smdx.Type != t {
		panic(fmt.Errorf("type mismatch: point=%s, actual=%s, expected=%s", p.smdx.Id, t, p.smdx.Type))
	}
	p.err = nil
	p.value = v
}

// The identifier of the point (relative to the block)
func (p *point) Id() string {
	return p.smdx.Id
}

// Answers an error if either the point itself or the
// scale factor point (if there is one) is in error.
func (p *point) Error() error {
	if p.err != nil {
		return p.err
	}
	if p.scaleFactor != nil {
		if err := p.scaleFactor.Error(); err != nil {
			return fmt.Errorf("error in scale factor point: %v", err)
		}
	}
	return nil
}

func (p *point) SetError(err error) {
	if err != nil {
		p.value = nil
		p.err = err
	}
}

// The type name of the point.
func (p *point) Type() string {
	return p.smdx.Type
}

func (p *point) Offset() uint16 {
	return p.smdx.Offset
}

func (p *point) Length() uint16 {
	if p.smdx.Length != 0 {
		return p.smdx.Length
	} else {
		return typelen.Length(p.smdx.Type)
	}
}

// NotImplemented validates if the actual value matches any of the
// unimplemented values according to the specification. In this case
// it will return true. NotImplemented can be used to check if the
// results returned by any of the getter functions is valid.
func (p *point) NotImplemented() bool {
	p.checkerror()

	notImplemented := false
	switch v := p.value.(type) {
	case float32:
		if 0x7fc00000 == math.Float32bits(v) {
			notImplemented = true
		}
	case int16:
		if v == int16(math.MinInt16) {
			notImplemented = true
		}
	case int32:
		if v == int32(math.MinInt32) {
			notImplemented = true
		}
	case int64:
		if v == int64(math.MinInt64) {
			notImplemented = true
		}
	case uint16:
		if v == uint16(math.MaxUint16) {
			notImplemented = true
		}
	case uint32:
		if v == uint32(math.MaxUint32) {
			notImplemented = true
		}
	case uint64:
		if v == uint64(math.MaxUint64) {
			notImplemented = true
		}
	case sunspec.Acc16:
		if v == sunspec.Acc16(math.MaxUint16) {
			notImplemented = true
		}
	case sunspec.Acc32:
		if v == sunspec.Acc32(math.MaxUint32) {
			notImplemented = true
		}
	case sunspec.Acc64:
		if v == sunspec.Acc64(math.MaxUint64) {
			notImplemented = true
		}
	case sunspec.Enum16:
		if v == sunspec.Enum16(math.MaxUint16) {
			notImplemented = true
		}
	case sunspec.Enum32:
		if v == sunspec.Enum32(math.MaxUint32) {
			notImplemented = true
		}
	case sunspec.Bitfield16:
		if v == sunspec.Bitfield16(math.MaxUint16) {
			notImplemented = true
		}
	case sunspec.Bitfield32:
		if v == sunspec.Bitfield32(math.MaxUint32) {
			notImplemented = true
		}
	case sunspec.Ipaddr:
		// 4 bytes
		if binary.BigEndian.Uint32(v[0:]) == 0 {
			notImplemented = true
		}
	case sunspec.Ipv6addr:
		// 2x8 bytes
		if bytes.Equal(v[:], []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}) {
			notImplemented = true
		}
	case sunspec.Eui48:
		// 6 bytes starting at 2
		if bytes.Equal(v[2:], []byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}) {
			notImplemented = true
		}
	case sunspec.ScaleFactor:
		if v == sunspec.ScaleFactor(math.MinInt16) {
			notImplemented = true
		}
	}

	return notImplemented
}

// Scales the point value with the associated scaling factor,
// if any, returning a float64 value.
func (p *point) ScaledValue() float64 {
	p.checkerror()
	var f float64
	switch v := p.value.(type) {
	case sunspec.Acc16:
		f = float64(v)
	case sunspec.Acc32:
		f = float64(v)
	case sunspec.Acc64:
		f = float64(v)
	case float32:
		f = float64(v)
	case int16:
		f = float64(v)
	case int32:
		f = float64(v)
	case int64:
		f = float64(v)
	case uint16:
		f = float64(v)
	case uint32:
		f = float64(v)
	case uint64:
		f = float64(v)
	default:
		panic(fmt.Errorf("attempt to use non-numeric value as scaled value: point=%s", p.Id()))
	}

	// Unimplemented value
	if p.NotImplemented() {
		return math.NaN()
	}

	sf := sunspec.ScaleFactor(0)
	if p.scaleFactor != nil {
		sf = p.scaleFactor.ScaleFactor()
	} else if p.smdx.ScaleFactor != "" {
		if v, err := strconv.Atoi(p.smdx.ScaleFactor); err != nil {
			sf = sunspec.ScaleFactor(v)
		}
	}

	// Unimplemented scale factor value
	if int16(sf) == int16(math.MinInt16) {
		return math.NaN()
	}

	return f * math.Pow(10, float64(sf))
}

// The value of the point, without regard to its actual type.
func (p *point) Value() interface{} {
	p.checkerror()
	if p.NotImplemented() {
		return sunspec.NotImplemented(struct{}{})
	}
	return p.value
}

// Set the value of the point, checking its type first.
func (p *point) SetValue(v interface{}) (result error) {

	defer func() {
		// use of recover here avoids the need to check
		// the type twice.
		if e := recover(); e != nil {
			if err, ok := e.(error); ok {
				result = err
			} else {
				panic(e)
			}
		}
	}()

	if v == nil {
		p.value = nil
		p.err = errNotInitialized
		if p.Type() == typelabel.ScaleFactor {
			p.block.invalidate(p)
		}
		return nil
	} else {
		switch t := v.(type) {
		case sunspec.Acc16:
			p.SetAcc16(t)
		case sunspec.Acc32:
			p.SetAcc32(t)
		case sunspec.Acc64:
			p.SetAcc64(t)
		case sunspec.Bitfield16:
			p.SetBitfield16(t)
		case sunspec.Bitfield32:
			p.SetBitfield32(t)
		case sunspec.Count:
			p.SetCount(t)
		case sunspec.Enum16:
			p.SetEnum16(t)
		case sunspec.Enum32:
			p.SetEnum32(t)
		case sunspec.Eui48:
			p.SetEui48(t)
		case float32:
			p.SetFloat32(t)
		case int16:
			p.SetInt16(t)
		case int32:
			p.SetInt32(t)
		case int64:
			p.SetInt64(t)
		case sunspec.Ipaddr:
			p.SetIpaddr(t)
		case sunspec.Ipv6addr:
			p.SetIpv6addr(t)
		case sunspec.Pad:
			p.SetPad(t)
		case string:
			p.SetStringValue(t)
		case sunspec.ScaleFactor:
			p.SetScaleFactor(t)
		case uint16:
			p.SetUint16(t)
		case uint32:
			p.SetUint32(t)
		case uint64:
			p.SetUint64(t)
		default:
			return fmt.Errorf("bad value type for point: %s: %v", p.Id(), v)
		}
	}
	return nil
}

// Strongly typed accessor methods

func (p *point) Acc16() sunspec.Acc16 {
	return p.checkerror().(sunspec.Acc16)
}

func (p *point) Acc32() sunspec.Acc32 {
	return p.checkerror().(sunspec.Acc32)
}

func (p *point) Acc64() sunspec.Acc64 {
	return p.checkerror().(sunspec.Acc64)
}

func (p *point) Bitfield16() sunspec.Bitfield16 {
	return p.checkerror().(sunspec.Bitfield16)
}

func (p *point) Bitfield32() sunspec.Bitfield32 {
	return p.checkerror().(sunspec.Bitfield32)
}

func (p *point) Count() sunspec.Count {
	return p.checkerror().(sunspec.Count)
}

func (p *point) Enum16() sunspec.Enum16 {
	return p.checkerror().(sunspec.Enum16)
}

func (p *point) Enum32() sunspec.Enum32 {
	return p.checkerror().(sunspec.Enum32)
}

func (p *point) Eui48() sunspec.Eui48 {
	return p.checkerror().(sunspec.Eui48)
}

func (p *point) Float32() float32 {
	return p.checkerror().(float32)
}

func (p *point) Int16() int16 {
	return p.checkerror().(int16)
}

func (p *point) Int32() int32 {
	return p.checkerror().(int32)
}

func (p *point) Int64() int64 {
	return p.checkerror().(int64)
}

func (p *point) Ipaddr() sunspec.Ipaddr {
	return p.checkerror().(sunspec.Ipaddr)
}

func (p *point) Ipv6addr() sunspec.Ipv6addr {
	return p.checkerror().(sunspec.Ipv6addr)
}

func (p *point) Pad() sunspec.Pad {
	return p.checkerror().(sunspec.Pad)
}

func (p *point) StringValue() string {
	return p.checkerror().(string)
}

func (p *point) ScaleFactor() sunspec.ScaleFactor {
	return p.checkerror().(sunspec.ScaleFactor)
}

func (p *point) Uint16() uint16 {
	return p.checkerror().(uint16)
}

func (p *point) Uint32() uint32 {
	return p.checkerror().(uint32)
}

func (p *point) Uint64() uint64 {
	return p.checkerror().(uint64)
}

func (p *point) SetAcc16(v sunspec.Acc16) {
	p.checktype(typelabel.Acc16, v)
}

func (p *point) SetAcc32(v sunspec.Acc32) {
	p.checktype(typelabel.Acc32, v)
}

func (p *point) SetAcc64(v sunspec.Acc64) {
	p.checktype(typelabel.Acc64, v)
}

func (p *point) SetBitfield16(v sunspec.Bitfield16) {
	p.checktype(typelabel.Bitfield16, v)
}

func (p *point) SetBitfield32(v sunspec.Bitfield32) {
	p.checktype(typelabel.Bitfield32, v)
}

func (p *point) SetCount(v sunspec.Count) {
	p.checktype(typelabel.Count, v)
}

func (p *point) SetEnum16(v sunspec.Enum16) {
	p.checktype(typelabel.Enum16, v)
}

func (p *point) SetEnum32(v sunspec.Enum32) {
	p.checktype(typelabel.Enum32, v)
}

func (p *point) SetEui48(v sunspec.Eui48) {
	p.checktype(typelabel.Eui48, v)
}

func (p *point) SetFloat32(v float32) {
	p.checktype(typelabel.Float32, v)
}

func (p *point) SetInt16(v int16) {
	p.checktype(typelabel.Int16, v)
}

func (p *point) SetInt32(v int32) {
	p.checktype(typelabel.Int32, v)
}

func (p *point) SetInt64(v int64) {
	p.checktype(typelabel.Int64, v)
}

func (p *point) SetIpaddr(v sunspec.Ipaddr) {
	p.checktype(typelabel.Ipaddr, v)
}

func (p *point) SetIpv6addr(v sunspec.Ipv6addr) {
	p.checktype(typelabel.Ipv6addr, v)
}

func (p *point) SetPad(v sunspec.Pad) {
	p.checktype(typelabel.Pad, v)
}

func (p *point) SetStringValue(v string) {
	p.checktype(typelabel.String, v)
}

func (p *point) SetScaleFactor(v sunspec.ScaleFactor) {
	p.checktype(typelabel.ScaleFactor, v)
	p.block.invalidate(p)
}

func (p *point) SetUint16(v uint16) {
	p.checktype(typelabel.Uint16, v)
}

func (p *point) SetUint32(v uint32) {
	p.checktype(typelabel.Uint32, v)
}

func (p *point) SetUint64(v uint64) {
	p.checktype(typelabel.Uint64, v)
}

func (p *point) Marshal(bytes []byte) error {
	if p.value == nil {
		return nil
	}
	switch p.Type() {
	case typelabel.Acc16:
		binary.BigEndian.PutUint16(bytes, uint16(p.value.(sunspec.Acc16)))
	case typelabel.Acc32:
		binary.BigEndian.PutUint32(bytes, uint32(p.value.(sunspec.Acc32)))
	case typelabel.Acc64:
		binary.BigEndian.PutUint64(bytes, uint64(p.value.(sunspec.Acc64)))
	case typelabel.Bitfield16:
		binary.BigEndian.PutUint16(bytes, uint16(p.value.(sunspec.Bitfield16)))
	case typelabel.Bitfield32:
		binary.BigEndian.PutUint32(bytes, uint32(p.value.(sunspec.Bitfield32)))
	case typelabel.Count:
		binary.BigEndian.PutUint16(bytes, uint16(p.value.(sunspec.Count)))
	case typelabel.Enum16:
		binary.BigEndian.PutUint16(bytes, uint16(p.value.(sunspec.Enum16)))
	case typelabel.Enum32:
		binary.BigEndian.PutUint32(bytes, uint32(p.value.(sunspec.Enum32)))
	case typelabel.Eui48:
		eui := p.value.(sunspec.Eui48)
		copy(bytes, eui[0:])
	case typelabel.Float32:
		binary.BigEndian.PutUint32(bytes, math.Float32bits(p.value.(float32)))
	case typelabel.Int16:
		binary.BigEndian.PutUint16(bytes, uint16(p.value.(int16)))
	case typelabel.Int32:
		binary.BigEndian.PutUint32(bytes, uint32(p.value.(int32)))
	case typelabel.Int64:
		binary.BigEndian.PutUint64(bytes, uint64(p.value.(int64)))
	case typelabel.Ipaddr:
		ip := p.value.(sunspec.Ipaddr)
		copy(bytes, ip[0:])
	case typelabel.Ipv6addr:
		ip := p.value.(sunspec.Ipv6addr)
		copy(bytes, ip[0:])
	case typelabel.Pad:
		binary.BigEndian.PutUint16(bytes, uint16(p.value.(sunspec.Pad)))
	case typelabel.String:
		s := p.value.(string)
		l := int(p.Length() * 2)
		copy(bytes, s[0:])
		for i := len(s); i < l; i++ {
			bytes[i] = 0
		}
	case typelabel.ScaleFactor:
		binary.BigEndian.PutUint16(bytes, uint16(p.value.(sunspec.ScaleFactor)))
	case typelabel.Uint16:
		binary.BigEndian.PutUint16(bytes, p.value.(uint16))
	case typelabel.Uint32:
		binary.BigEndian.PutUint32(bytes, p.value.(uint32))
	case typelabel.Uint64:
		binary.BigEndian.PutUint64(bytes, p.value.(uint64))
	default:
		return errBadType
	}
	return nil
}

func (p *point) Unmarshal(bytes []byte) error {
	if len(bytes) < int(p.Length()*2) {
		return errLength
	}
	switch p.Type() {
	case typelabel.Acc16:
		p.SetAcc16(sunspec.Acc16(binary.BigEndian.Uint16(bytes)))
	case typelabel.Acc32:
		p.SetAcc32(sunspec.Acc32(binary.BigEndian.Uint32(bytes)))
	case typelabel.Acc64:
		p.SetAcc64(sunspec.Acc64(binary.BigEndian.Uint64(bytes)))
	case typelabel.Bitfield16:
		p.SetBitfield16(sunspec.Bitfield16(binary.BigEndian.Uint16(bytes)))
	case typelabel.Bitfield32:
		p.SetBitfield32(sunspec.Bitfield32(binary.BigEndian.Uint32(bytes)))
	case typelabel.Count:
		p.SetCount(sunspec.Count(binary.BigEndian.Uint16(bytes)))
	case typelabel.Enum16:
		p.SetEnum16(sunspec.Enum16(binary.BigEndian.Uint16(bytes)))
	case typelabel.Enum32:
		p.SetEnum32(sunspec.Enum32(binary.BigEndian.Uint32(bytes)))
	case typelabel.Eui48:
		var eui sunspec.Eui48
		copy(eui[0:], bytes)
		p.SetEui48(eui)
	case typelabel.Float32:
		u32 := binary.BigEndian.Uint32(bytes)
		p.SetFloat32(math.Float32frombits(u32))
	case typelabel.Int16:
		p.SetInt16(int16(binary.BigEndian.Uint16(bytes)))
	case typelabel.Int32:
		p.SetInt32(int32(binary.BigEndian.Uint32(bytes)))
	case typelabel.Int64:
		p.SetInt64(int64(binary.BigEndian.Uint64(bytes)))
	case typelabel.Ipaddr:
		var ip sunspec.Ipaddr
		copy(ip[0:], bytes)
		p.SetIpaddr(ip)
	case typelabel.Ipv6addr:
		var ip sunspec.Ipv6addr
		copy(ip[0:], bytes)
		p.SetIpv6addr(ip)
	case typelabel.Pad:
		p.SetPad(sunspec.Pad(binary.BigEndian.Uint16(bytes)))
	case typelabel.String:
		var b = make([]byte, p.Length(), p.Length())
		copy(b[0:], bytes)
		for i, y := range b {
			if y == 0 {
				b = b[0:i]
				break
			}
		}
		p.SetStringValue(string(b))
	case typelabel.ScaleFactor:
		p.SetScaleFactor(sunspec.ScaleFactor(binary.BigEndian.Uint16(bytes)))
	case typelabel.Uint16:
		p.SetUint16(uint16(binary.BigEndian.Uint16(bytes)))
	case typelabel.Uint32:
		p.SetUint32(uint32(binary.BigEndian.Uint32(bytes)))
	case typelabel.Uint64:
		p.SetUint64(uint64(binary.BigEndian.Uint64(bytes)))
	default:
		return errBadType
	}
	return nil
}

func (p *point) ScaleFactorPoint() spi.PointSPI {
	if p.scaleFactor == nil {
		return nil
	} else {
		return p.scaleFactor.(spi.PointSPI)
	}
}

func (p *point) MarshalXML() string {
	if p.err != nil {
		panic(p.err)
	}
	switch p.Type() {
	case typelabel.Bitfield16:
		return p.Bitfield16().String()
	case typelabel.Bitfield32:
		return p.Bitfield32().String()
	case typelabel.Pad:
		return p.Pad().String()
	case typelabel.Ipaddr:
		return p.Ipaddr().String()
	case typelabel.Ipv6addr:
		return p.Ipv6addr().String()
	case typelabel.Eui48:
		return p.Eui48().String()
	default:
		return fmt.Sprintf("%v", p.value)
	}
}

func (p *point) UnmarshalXML(s string) error {
	checkzero := func(buf [8]byte, n int) error {
		b0 := buf[0]
		for _, b := range buf[0 : 8-n] {
			if b != b0 {
				return errValueTooLarge
			}
		}
		return nil
	}
	intApply := func(n int) error {
		if v, e := strconv.Atoi(s); e != nil {
			return e
		} else {
			buf := [8]byte{}
			binary.BigEndian.PutUint64(buf[0:], uint64(v))
			if err := checkzero(buf, n); err != nil {
				return err
			} else {
				return p.Unmarshal(buf[8-n:])
			}
		}
	}
	uintApply := func(n int) error {
		if v, e := strconv.ParseUint(s, 10, 64); e != nil {
			return e
		} else {
			buf := [8]byte{}
			binary.BigEndian.PutUint64(buf[0:], v)
			if err := checkzero(buf, n); err != nil {
				return err
			} else {
				return p.Unmarshal(buf[8-n:])
			}
		}
	}

	xintApply := func(n int) error {
		var x uint64
		fmt.Sscanf(s, "0x%x", &x)
		buf := [8]byte{}
		binary.BigEndian.PutUint64(buf[0:], x)
		if err := checkzero(buf, n); err != nil {
			return err
		} else {
			return p.Unmarshal(buf[8-n:])
		}
	}

	switch p.Type() {
	case typelabel.Bitfield16, typelabel.Pad:
		xintApply(2)
	case typelabel.Bitfield32:
		xintApply(4)
	case typelabel.Ipaddr:
		var buf sunspec.Ipaddr
		if _, err := fmt.Sscanf(s, "%d.%d.%d.%d", &buf[0], &buf[1], &buf[2], &buf[3]); err != nil {
			return err
		} else {
			p.SetIpaddr(buf)
		}
	case typelabel.Ipv6addr:
		var ubuf [8]uint16
		if _, err := fmt.Sscanf(s, "%04x:%04x:%04x:%04x:%04x:%04x:%04x:%04x", &ubuf[0], &ubuf[1], &ubuf[2], &ubuf[3], &ubuf[4], &ubuf[5], &ubuf[6], &ubuf[7]); err != nil {
			return err
		} else {
			var bbuf [16]byte
			for i, u := range ubuf {
				binary.BigEndian.PutUint16(bbuf[i*2:i*2+2], u)
			}
			return p.Unmarshal(bbuf[0:])
		}
	case typelabel.Eui48:
		var buf sunspec.Eui48
		// first word is unused - ignore
		if _, err := fmt.Sscanf(s, "%02x:%02x:%02x:%02x:%02x:%02x", &buf[2], &buf[3], &buf[4], &buf[5], &buf[6], &buf[7]); err != nil {
			return err
		} else {
			p.SetEui48(buf)
		}
	case typelabel.String:
		p.SetStringValue(s)
	case typelabel.Int16:
		return intApply(2)
	case typelabel.ScaleFactor:
		return intApply(2)
	case typelabel.Int32:
		return intApply(4)
	case typelabel.Int64:
		return intApply(8)
	case typelabel.Enum16:
		return uintApply(2)
	case typelabel.Count:
		return uintApply(2)
	case typelabel.Enum32:
		return uintApply(4)
	case typelabel.Acc16:
		return uintApply(2)
	case typelabel.Acc32:
		return uintApply(4)
	case typelabel.Acc64:
		return uintApply(8)
	case typelabel.Uint16:
		return uintApply(2)
	case typelabel.Uint32:
		return uintApply(4)
	case typelabel.Uint64:
		return uintApply(8)
	case typelabel.Float32:
		if f, err := strconv.ParseFloat(s, 32); err != nil {
			return err
		} else {
			p.SetFloat32(float32(f))
		}
	default:
		panic(errBadType)
	}
	return nil
}
