package model63001

// Code generated by ../../generators/models.go. DO NOT EDIT.
// You can regenerate it by running 'go generate ./models' from the directory above.

import (
	"github.com/andig/gosunspec"
	"github.com/andig/gosunspec/smdx"
	"github.com/andig/gosunspec/typelabel"
)

// Block63001 - SunSpec Test Model 1 -

const (
	ModelID          = 63001
	ModelLabel       = "SunSpec Test Model 1"
	ModelDescription = ""
)

const (
	Acc16        = "acc16"
	Acc16_u      = "acc16_u"
	Acc32        = "acc32"
	Acc32_u      = "acc32_u"
	Acc64        = "acc64"
	Acc64_u      = "acc64_u"
	Bitfield16   = "bitfield16"
	Bitfield16_u = "bitfield16_u"
	Bitfield32   = "bitfield32"
	Bitfield32_u = "bitfield32_u"
	Enum16       = "enum16"
	Enum16_u     = "enum16_u"
	Enum32       = "enum32"
	Enum32_u     = "enum32_u"
	Float32      = "float32"
	Float32_u    = "float32_u"
	Int16_1      = "int16_1"
	Int16_11     = "int16_11"
	Int16_12     = "int16_12"
	Int16_2      = "int16_2"
	Int16_3      = "int16_3"
	Int16_4      = "int16_4"
	Int16_5      = "int16_5"
	Int16_u      = "int16_u"
	Int32        = "int32"
	Int32_1      = "int32_1"
	Int32_2      = "int32_2"
	Int32_3      = "int32_3"
	Int32_4      = "int32_4"
	Int32_5      = "int32_5"
	Int32_u      = "int32_u"
	Int64        = "int64"
	Int64_u      = "int64_u"
	Ipaddr       = "ipaddr"
	Ipaddr_u     = "ipaddr_u"
	Ipv6addr     = "ipv6addr"
	Ipv6addr_u   = "ipv6addr_u"
	Pad_1        = "pad_1"
	Pad_2        = "pad_2"
	String       = "string"
	String_u     = "string_u"
	Sunssf_1     = "sunssf_1"
	Sunssf_2     = "sunssf_2"
	Sunssf_3     = "sunssf_3"
	Sunssf_4     = "sunssf_4"
	Sunssf_5     = "sunssf_5"
	Sunssf_6     = "sunssf_6"
	Sunssf_7     = "sunssf_7"
	Sunssf_8     = "sunssf_8"
	Sunssf_9     = "sunssf_9"
	Uint16_1     = "uint16_1"
	Uint16_11    = "uint16_11"
	Uint16_12    = "uint16_12"
	Uint16_13    = "uint16_13"
	Uint16_2     = "uint16_2"
	Uint16_3     = "uint16_3"
	Uint16_4     = "uint16_4"
	Uint16_5     = "uint16_5"
	Uint16_u     = "uint16_u"
	Uint32       = "uint32"
	Uint32_1     = "uint32_1"
	Uint32_2     = "uint32_2"
	Uint32_3     = "uint32_3"
	Uint32_4     = "uint32_4"
	Uint32_5     = "uint32_5"
	Uint32_u     = "uint32_u"
)

type Block63001Repeat struct {
	sunssf_8  sunspec.ScaleFactor `sunspec:"offset=0,len=1,access=r"`
	int16_11  int16               `sunspec:"offset=1,len=1,sf=sunssf_8,access=rw"`
	int16_12  int16               `sunspec:"offset=2,len=1,sf=sunssf_9,access=r"`
	int16_u   int16               `sunspec:"offset=3,len=1,access=r"`
	uint16_11 uint16              `sunspec:"offset=4,len=1,sf=sunssf_8,access=rw"`
	uint16_12 uint16              `sunspec:"offset=5,len=1,sf=sunssf_9,access=r"`
	uint16_13 uint16              `sunspec:"offset=6,len=1,access=r"`
	uint16_u  uint16              `sunspec:"offset=7,len=1,access=r"`
	int32     int32               `sunspec:"offset=8,len=2,sf=sunssf_1,access=rw"`
	int32_u   int32               `sunspec:"offset=10,len=2,access=r"`
	uint32    uint32              `sunspec:"offset=12,len=2,sf=sunssf_9,access=rw"`
	uint32_u  uint32              `sunspec:"offset=14,len=2,access=r"`
	sunssf_9  sunspec.ScaleFactor `sunspec:"offset=16,len=1,access=r"`
	pad_2     sunspec.Pad         `sunspec:"offset=17,len=1,access=r"`
}

type Block63001 struct {
	sunssf_1     sunspec.ScaleFactor `sunspec:"offset=0,len=1,access=r"`
	sunssf_2     sunspec.ScaleFactor `sunspec:"offset=1,len=1,access=r"`
	sunssf_3     sunspec.ScaleFactor `sunspec:"offset=2,len=1,access=r"`
	sunssf_4     sunspec.ScaleFactor `sunspec:"offset=3,len=1,access=r"`
	int16_1      int16               `sunspec:"offset=4,len=1,sf=sunssf_1,access=r"`
	int16_2      int16               `sunspec:"offset=5,len=1,sf=sunssf_2,access=r"`
	int16_3      int16               `sunspec:"offset=6,len=1,sf=sunssf_3,access=r"`
	int16_4      int16               `sunspec:"offset=7,len=1,sf=sunssf_4,access=rw"`
	int16_5      int16               `sunspec:"offset=8,len=1,access=r"`
	int16_u      int16               `sunspec:"offset=9,len=1,access=r"`
	uint16_1     uint16              `sunspec:"offset=10,len=1,sf=sunssf_1,access=r"`
	uint16_2     uint16              `sunspec:"offset=11,len=1,sf=sunssf_2,access=r"`
	uint16_3     uint16              `sunspec:"offset=12,len=1,sf=sunssf_3,access=r"`
	uint16_4     uint16              `sunspec:"offset=13,len=1,sf=sunssf_4,access=rw"`
	uint16_5     uint16              `sunspec:"offset=14,len=1,access=r"`
	uint16_u     uint16              `sunspec:"offset=15,len=1,access=r"`
	acc16        sunspec.Acc16       `sunspec:"offset=16,len=1,access=r"`
	acc16_u      sunspec.Acc16       `sunspec:"offset=17,len=1,access=r"`
	enum16       sunspec.Enum16      `sunspec:"offset=18,len=1,access=r"`
	enum16_u     sunspec.Enum16      `sunspec:"offset=19,len=1,access=r"`
	bitfield16   sunspec.Bitfield16  `sunspec:"offset=20,len=1,access=r"`
	bitfield16_u sunspec.Bitfield16  `sunspec:"offset=21,len=1,access=r"`
	int32_1      int32               `sunspec:"offset=22,len=2,sf=sunssf_5,access=r"`
	int32_2      int32               `sunspec:"offset=24,len=2,sf=sunssf_6,access=r"`
	int32_3      int32               `sunspec:"offset=26,len=2,sf=sunssf_7,access=rw"`
	int32_4      int32               `sunspec:"offset=28,len=2,access=r"`
	int32_5      int32               `sunspec:"offset=30,len=2,access=r"`
	int32_u      int32               `sunspec:"offset=32,len=2,access=r"`
	uint32_1     uint32              `sunspec:"offset=34,len=2,sf=sunssf_5,access=r"`
	uint32_2     uint32              `sunspec:"offset=36,len=2,sf=sunssf_6,access=r"`
	uint32_3     uint32              `sunspec:"offset=38,len=2,sf=sunssf_7,access=rw"`
	uint32_4     uint32              `sunspec:"offset=40,len=2,sf=1,access=r"`
	uint32_5     uint32              `sunspec:"offset=42,len=2,access=r"`
	uint32_u     uint32              `sunspec:"offset=44,len=2,access=r"`
	acc32        sunspec.Acc32       `sunspec:"offset=46,len=2,access=r"`
	acc32_u      sunspec.Acc32       `sunspec:"offset=48,len=2,access=r"`
	enum32       sunspec.Enum32      `sunspec:"offset=50,len=2,access=r"`
	enum32_u     sunspec.Enum32      `sunspec:"offset=52,len=2,access=r"`
	bitfield32   sunspec.Bitfield32  `sunspec:"offset=54,len=2,access=r"`
	bitfield32_u sunspec.Bitfield32  `sunspec:"offset=56,len=2,access=r"`
	ipaddr       sunspec.Ipaddr      `sunspec:"offset=58,len=2,access=rw"`
	ipaddr_u     sunspec.Ipaddr      `sunspec:"offset=60,len=2,access=r"`
	int64        int64               `sunspec:"offset=62,len=4,access=rw"`
	int64_u      int64               `sunspec:"offset=66,len=4,access=r"`
	acc64        sunspec.Acc64       `sunspec:"offset=70,len=4,access=r"`
	acc64_u      sunspec.Acc64       `sunspec:"offset=74,len=4,access=r"`
	ipv6addr     sunspec.Ipv6addr    `sunspec:"offset=78,len=8,access=r"`
	ipv6addr_u   sunspec.Ipv6addr    `sunspec:"offset=86,len=8,access=r"`
	float32      float32             `sunspec:"offset=94,len=2,access=rw"`
	float32_u    float32             `sunspec:"offset=96,len=2,access=r"`
	string       string              `sunspec:"offset=98,len=16,access=rw"`
	string_u     string              `sunspec:"offset=114,len=16,access=r"`
	sunssf_5     sunspec.ScaleFactor `sunspec:"offset=130,len=1,access=r"`
	sunssf_6     sunspec.ScaleFactor `sunspec:"offset=131,len=1,access=r"`
	sunssf_7     sunspec.ScaleFactor `sunspec:"offset=132,len=1,access=r"`
	pad_1        sunspec.Pad         `sunspec:"offset=133,len=1,access=r"`

	Repeats []Block63001Repeat
}

func (block *Block63001) GetId() sunspec.ModelId {
	return ModelID
}

func init() {
	smdx.RegisterModel(&smdx.ModelElement{
		Id:     ModelID,
		Name:   "model_63001",
		Length: 152,
		Blocks: []smdx.BlockElement{
			{
				Length: 134,
				Points: []smdx.PointElement{
					{Id: Sunssf_1, Offset: 0, Type: typelabel.ScaleFactor, Access: "r", Length: 1},
					{Id: Sunssf_2, Offset: 1, Type: typelabel.ScaleFactor, Access: "r", Length: 1},
					{Id: Sunssf_3, Offset: 2, Type: typelabel.ScaleFactor, Access: "r", Length: 1},
					{Id: Sunssf_4, Offset: 3, Type: typelabel.ScaleFactor, Access: "r", Length: 1},
					{Id: Int16_1, Offset: 4, Type: typelabel.Int16, ScaleFactor: "sunssf_1", Access: "r", Length: 1},
					{Id: Int16_2, Offset: 5, Type: typelabel.Int16, ScaleFactor: "sunssf_2", Access: "r", Length: 1},
					{Id: Int16_3, Offset: 6, Type: typelabel.Int16, ScaleFactor: "sunssf_3", Access: "r", Length: 1},
					{Id: Int16_4, Offset: 7, Type: typelabel.Int16, ScaleFactor: "sunssf_4", Access: "rw", Length: 1},
					{Id: Int16_5, Offset: 8, Type: typelabel.Int16, Access: "r", Length: 1},
					{Id: Int16_u, Offset: 9, Type: typelabel.Int16, Access: "r", Length: 1},
					{Id: Uint16_1, Offset: 10, Type: typelabel.Uint16, ScaleFactor: "sunssf_1", Access: "r", Length: 1},
					{Id: Uint16_2, Offset: 11, Type: typelabel.Uint16, ScaleFactor: "sunssf_2", Access: "r", Length: 1},
					{Id: Uint16_3, Offset: 12, Type: typelabel.Uint16, ScaleFactor: "sunssf_3", Access: "r", Length: 1},
					{Id: Uint16_4, Offset: 13, Type: typelabel.Uint16, ScaleFactor: "sunssf_4", Access: "rw", Length: 1},
					{Id: Uint16_5, Offset: 14, Type: typelabel.Uint16, Access: "r", Length: 1},
					{Id: Uint16_u, Offset: 15, Type: typelabel.Uint16, Access: "r", Length: 1},
					{Id: Acc16, Offset: 16, Type: typelabel.Acc16, Access: "r", Length: 1},
					{Id: Acc16_u, Offset: 17, Type: typelabel.Acc16, Access: "r", Length: 1},
					{Id: Enum16, Offset: 18, Type: typelabel.Enum16, Access: "r", Length: 1},
					{Id: Enum16_u, Offset: 19, Type: typelabel.Enum16, Access: "r", Length: 1},
					{Id: Bitfield16, Offset: 20, Type: typelabel.Bitfield16, Access: "r", Length: 1},
					{Id: Bitfield16_u, Offset: 21, Type: typelabel.Bitfield16, Access: "r", Length: 1},
					{Id: Int32_1, Offset: 22, Type: typelabel.Int32, ScaleFactor: "sunssf_5", Access: "r", Length: 2},
					{Id: Int32_2, Offset: 24, Type: typelabel.Int32, ScaleFactor: "sunssf_6", Access: "r", Length: 2},
					{Id: Int32_3, Offset: 26, Type: typelabel.Int32, ScaleFactor: "sunssf_7", Access: "rw", Length: 2},
					{Id: Int32_4, Offset: 28, Type: typelabel.Int32, Access: "r", Length: 2},
					{Id: Int32_5, Offset: 30, Type: typelabel.Int32, Access: "r", Length: 2},
					{Id: Int32_u, Offset: 32, Type: typelabel.Int32, Access: "r", Length: 2},
					{Id: Uint32_1, Offset: 34, Type: typelabel.Uint32, ScaleFactor: "sunssf_5", Access: "r", Length: 2},
					{Id: Uint32_2, Offset: 36, Type: typelabel.Uint32, ScaleFactor: "sunssf_6", Access: "r", Length: 2},
					{Id: Uint32_3, Offset: 38, Type: typelabel.Uint32, ScaleFactor: "sunssf_7", Access: "rw", Length: 2},
					{Id: Uint32_4, Offset: 40, Type: typelabel.Uint32, ScaleFactor: "1", Access: "r", Length: 2},
					{Id: Uint32_5, Offset: 42, Type: typelabel.Uint32, Access: "r", Length: 2},
					{Id: Uint32_u, Offset: 44, Type: typelabel.Uint32, Access: "r", Length: 2},
					{Id: Acc32, Offset: 46, Type: typelabel.Acc32, Access: "r", Length: 2},
					{Id: Acc32_u, Offset: 48, Type: typelabel.Acc32, Access: "r", Length: 2},
					{Id: Enum32, Offset: 50, Type: typelabel.Enum32, Access: "r", Length: 2},
					{Id: Enum32_u, Offset: 52, Type: typelabel.Enum32, Access: "r", Length: 2},
					{Id: Bitfield32, Offset: 54, Type: typelabel.Bitfield32, Access: "r", Length: 2},
					{Id: Bitfield32_u, Offset: 56, Type: typelabel.Bitfield32, Access: "r", Length: 2},
					{Id: Ipaddr, Offset: 58, Type: typelabel.Ipaddr, Access: "rw", Length: 2},
					{Id: Ipaddr_u, Offset: 60, Type: typelabel.Ipaddr, Access: "r", Length: 2},
					{Id: Int64, Offset: 62, Type: typelabel.Int64, Access: "rw", Length: 4},
					{Id: Int64_u, Offset: 66, Type: typelabel.Int64, Access: "r", Length: 4},
					{Id: Acc64, Offset: 70, Type: typelabel.Acc64, Access: "r", Length: 4},
					{Id: Acc64_u, Offset: 74, Type: typelabel.Acc64, Access: "r", Length: 4},
					{Id: Ipv6addr, Offset: 78, Type: typelabel.Ipv6addr, Access: "r", Length: 8},
					{Id: Ipv6addr_u, Offset: 86, Type: typelabel.Ipv6addr, Access: "r", Length: 8},
					{Id: Float32, Offset: 94, Type: typelabel.Float32, Access: "rw", Length: 2},
					{Id: Float32_u, Offset: 96, Type: typelabel.Float32, Access: "r", Length: 2},
					{Id: String, Offset: 98, Type: typelabel.String, Access: "rw", Length: 16},
					{Id: String_u, Offset: 114, Type: typelabel.String, Access: "r", Length: 16},
					{Id: Sunssf_5, Offset: 130, Type: typelabel.ScaleFactor, Access: "r", Length: 1},
					{Id: Sunssf_6, Offset: 131, Type: typelabel.ScaleFactor, Access: "r", Length: 1},
					{Id: Sunssf_7, Offset: 132, Type: typelabel.ScaleFactor, Access: "r", Length: 1},
					{Id: Pad_1, Offset: 133, Type: typelabel.Pad, Access: "r", Length: 1},
				},
			},
			{Name: "repeating",
				Length: 18,
				Type:   "repeating",
				Points: []smdx.PointElement{
					{Id: Sunssf_8, Offset: 0, Type: typelabel.ScaleFactor, Access: "r", Length: 1},
					{Id: Int16_11, Offset: 1, Type: typelabel.Int16, ScaleFactor: "sunssf_8", Access: "rw", Length: 1},
					{Id: Int16_12, Offset: 2, Type: typelabel.Int16, ScaleFactor: "sunssf_9", Access: "r", Length: 1},
					{Id: Int16_u, Offset: 3, Type: typelabel.Int16, Access: "r", Length: 1},
					{Id: Uint16_11, Offset: 4, Type: typelabel.Uint16, ScaleFactor: "sunssf_8", Access: "rw", Length: 1},
					{Id: Uint16_12, Offset: 5, Type: typelabel.Uint16, ScaleFactor: "sunssf_9", Access: "r", Length: 1},
					{Id: Uint16_13, Offset: 6, Type: typelabel.Uint16, Access: "r", Length: 1},
					{Id: Uint16_u, Offset: 7, Type: typelabel.Uint16, Access: "r", Length: 1},
					{Id: Int32, Offset: 8, Type: typelabel.Int32, ScaleFactor: "sunssf_1", Access: "rw", Length: 2},
					{Id: Int32_u, Offset: 10, Type: typelabel.Int32, Access: "r", Length: 2},
					{Id: Uint32, Offset: 12, Type: typelabel.Uint32, ScaleFactor: "sunssf_9", Access: "rw", Length: 2},
					{Id: Uint32_u, Offset: 14, Type: typelabel.Uint32, Access: "r", Length: 2},
					{Id: Sunssf_9, Offset: 16, Type: typelabel.ScaleFactor, Access: "r", Length: 1},
					{Id: Pad_2, Offset: 17, Type: typelabel.Pad, Access: "r", Length: 1},
				},
			},
		}})
}
