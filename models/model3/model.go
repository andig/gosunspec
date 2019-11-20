// NOTICE
// This file was automatically generated by ../../generators/models.go. Do not edit it!
// You can regenerate it by running 'go generate ./models' from the directory above.

package model3

import (
	"github.com/andig/gosunspec"
	"github.com/andig/gosunspec/smdx"
	"github.com/andig/gosunspec/typelabel"
)

// Block3 - Secure Dataset Read Request - Request a digital signature over a specified set of data registers

const (
	ModelID          = 3
	ModelLabel       = "Secure Dataset Read Request"
	ModelDescription = "Request a digital signature over a specified set of data registers"
)

const (
	Alg   = "Alg"
	DS    = "DS"
	Ms    = "Ms"
	N     = "N"
	Off1  = "Off1"
	Off10 = "Off10"
	Off11 = "Off11"
	Off12 = "Off12"
	Off13 = "Off13"
	Off14 = "Off14"
	Off15 = "Off15"
	Off16 = "Off16"
	Off17 = "Off17"
	Off18 = "Off18"
	Off19 = "Off19"
	Off2  = "Off2"
	Off20 = "Off20"
	Off21 = "Off21"
	Off22 = "Off22"
	Off23 = "Off23"
	Off24 = "Off24"
	Off25 = "Off25"
	Off26 = "Off26"
	Off27 = "Off27"
	Off28 = "Off28"
	Off29 = "Off29"
	Off3  = "Off3"
	Off30 = "Off30"
	Off31 = "Off31"
	Off32 = "Off32"
	Off33 = "Off33"
	Off34 = "Off34"
	Off35 = "Off35"
	Off36 = "Off36"
	Off37 = "Off37"
	Off38 = "Off38"
	Off39 = "Off39"
	Off4  = "Off4"
	Off40 = "Off40"
	Off41 = "Off41"
	Off42 = "Off42"
	Off43 = "Off43"
	Off44 = "Off44"
	Off45 = "Off45"
	Off46 = "Off46"
	Off47 = "Off47"
	Off48 = "Off48"
	Off49 = "Off49"
	Off5  = "Off5"
	Off50 = "Off50"
	Off6  = "Off6"
	Off7  = "Off7"
	Off8  = "Off8"
	Off9  = "Off9"
	Role  = "Role"
	Seq   = "Seq"
	Ts    = "Ts"
	X     = "X"
)

type Block3Repeat struct {
	DS uint16 `sunspec:"offset=0,access=r"`
}

type Block3 struct {
	X     uint16         `sunspec:"offset=0,access=rw"`
	Off1  uint16         `sunspec:"offset=1,access=rw"`
	Off2  uint16         `sunspec:"offset=2,access=rw"`
	Off3  uint16         `sunspec:"offset=3,access=rw"`
	Off4  uint16         `sunspec:"offset=4,access=rw"`
	Off5  uint16         `sunspec:"offset=5,access=rw"`
	Off6  uint16         `sunspec:"offset=6,access=rw"`
	Off7  uint16         `sunspec:"offset=7,access=rw"`
	Off8  uint16         `sunspec:"offset=8,access=rw"`
	Off9  uint16         `sunspec:"offset=9,access=rw"`
	Off10 uint16         `sunspec:"offset=10,access=rw"`
	Off11 uint16         `sunspec:"offset=11,access=rw"`
	Off12 uint16         `sunspec:"offset=12,access=rw"`
	Off13 uint16         `sunspec:"offset=13,access=rw"`
	Off14 uint16         `sunspec:"offset=14,access=rw"`
	Off15 uint16         `sunspec:"offset=15,access=rw"`
	Off16 uint16         `sunspec:"offset=16,access=rw"`
	Off17 uint16         `sunspec:"offset=17,access=rw"`
	Off18 uint16         `sunspec:"offset=18,access=rw"`
	Off19 uint16         `sunspec:"offset=19,access=rw"`
	Off20 uint16         `sunspec:"offset=20,access=rw"`
	Off21 uint16         `sunspec:"offset=21,access=rw"`
	Off22 uint16         `sunspec:"offset=22,access=rw"`
	Off23 uint16         `sunspec:"offset=23,access=rw"`
	Off24 uint16         `sunspec:"offset=24,access=rw"`
	Off25 uint16         `sunspec:"offset=25,access=rw"`
	Off26 uint16         `sunspec:"offset=26,access=rw"`
	Off27 uint16         `sunspec:"offset=27,access=rw"`
	Off28 uint16         `sunspec:"offset=28,access=rw"`
	Off29 uint16         `sunspec:"offset=29,access=rw"`
	Off30 uint16         `sunspec:"offset=30,access=rw"`
	Off31 uint16         `sunspec:"offset=31,access=rw"`
	Off32 uint16         `sunspec:"offset=32,access=rw"`
	Off33 uint16         `sunspec:"offset=33,access=rw"`
	Off34 uint16         `sunspec:"offset=34,access=rw"`
	Off35 uint16         `sunspec:"offset=35,access=rw"`
	Off36 uint16         `sunspec:"offset=36,access=rw"`
	Off37 uint16         `sunspec:"offset=37,access=rw"`
	Off38 uint16         `sunspec:"offset=38,access=rw"`
	Off39 uint16         `sunspec:"offset=39,access=rw"`
	Off40 uint16         `sunspec:"offset=40,access=rw"`
	Off41 uint16         `sunspec:"offset=41,access=rw"`
	Off42 uint16         `sunspec:"offset=42,access=rw"`
	Off43 uint16         `sunspec:"offset=43,access=rw"`
	Off44 uint16         `sunspec:"offset=44,access=rw"`
	Off45 uint16         `sunspec:"offset=45,access=rw"`
	Off46 uint16         `sunspec:"offset=46,access=rw"`
	Off47 uint16         `sunspec:"offset=47,access=rw"`
	Off48 uint16         `sunspec:"offset=48,access=rw"`
	Off49 uint16         `sunspec:"offset=49,access=rw"`
	Off50 uint16         `sunspec:"offset=50,access=rw"`
	Ts    uint32         `sunspec:"offset=51,access=rw"`
	Ms    uint16         `sunspec:"offset=53,access=rw"`
	Seq   uint16         `sunspec:"offset=54,access=rw"`
	Role  uint16         `sunspec:"offset=55,access=rw"`
	Alg   sunspec.Enum16 `sunspec:"offset=56,access=r"`
	N     uint16         `sunspec:"offset=57,access=r"`

	Repeats []Block3Repeat
}

func (self *Block3) GetId() sunspec.ModelId {
	return ModelID
}

func init() {
	smdx.RegisterModel(&smdx.ModelElement{
		Id:     ModelID,
		Name:   "",
		Length: 59,
		Blocks: []smdx.BlockElement{
			smdx.BlockElement{
				Length: 58,

				Points: []smdx.PointElement{
					smdx.PointElement{Id: X, Offset: 0, Type: typelabel.Uint16, Access: "rw", Mandatory: true, Label: "X", Description: "Number of registers being requested"},
					smdx.PointElement{Id: Off1, Offset: 1, Type: typelabel.Uint16, Access: "rw", Mandatory: true, Label: "Offset1", Description: "Offset of value to read"},
					smdx.PointElement{Id: Off2, Offset: 2, Type: typelabel.Uint16, Access: "rw", Mandatory: true},
					smdx.PointElement{Id: Off3, Offset: 3, Type: typelabel.Uint16, Access: "rw", Mandatory: true},
					smdx.PointElement{Id: Off4, Offset: 4, Type: typelabel.Uint16, Access: "rw", Mandatory: true},
					smdx.PointElement{Id: Off5, Offset: 5, Type: typelabel.Uint16, Access: "rw", Mandatory: true},
					smdx.PointElement{Id: Off6, Offset: 6, Type: typelabel.Uint16, Access: "rw", Mandatory: true},
					smdx.PointElement{Id: Off7, Offset: 7, Type: typelabel.Uint16, Access: "rw", Mandatory: true},
					smdx.PointElement{Id: Off8, Offset: 8, Type: typelabel.Uint16, Access: "rw", Mandatory: true},
					smdx.PointElement{Id: Off9, Offset: 9, Type: typelabel.Uint16, Access: "rw", Mandatory: true},
					smdx.PointElement{Id: Off10, Offset: 10, Type: typelabel.Uint16, Access: "rw", Mandatory: true},
					smdx.PointElement{Id: Off11, Offset: 11, Type: typelabel.Uint16, Access: "rw", Mandatory: true},
					smdx.PointElement{Id: Off12, Offset: 12, Type: typelabel.Uint16, Access: "rw", Mandatory: true},
					smdx.PointElement{Id: Off13, Offset: 13, Type: typelabel.Uint16, Access: "rw", Mandatory: true},
					smdx.PointElement{Id: Off14, Offset: 14, Type: typelabel.Uint16, Access: "rw", Mandatory: true},
					smdx.PointElement{Id: Off15, Offset: 15, Type: typelabel.Uint16, Access: "rw", Mandatory: true},
					smdx.PointElement{Id: Off16, Offset: 16, Type: typelabel.Uint16, Access: "rw", Mandatory: true},
					smdx.PointElement{Id: Off17, Offset: 17, Type: typelabel.Uint16, Access: "rw", Mandatory: true},
					smdx.PointElement{Id: Off18, Offset: 18, Type: typelabel.Uint16, Access: "rw", Mandatory: true},
					smdx.PointElement{Id: Off19, Offset: 19, Type: typelabel.Uint16, Access: "rw", Mandatory: true},
					smdx.PointElement{Id: Off20, Offset: 20, Type: typelabel.Uint16, Access: "rw", Mandatory: true},
					smdx.PointElement{Id: Off21, Offset: 21, Type: typelabel.Uint16, Access: "rw", Mandatory: true},
					smdx.PointElement{Id: Off22, Offset: 22, Type: typelabel.Uint16, Access: "rw", Mandatory: true},
					smdx.PointElement{Id: Off23, Offset: 23, Type: typelabel.Uint16, Access: "rw", Mandatory: true},
					smdx.PointElement{Id: Off24, Offset: 24, Type: typelabel.Uint16, Access: "rw", Mandatory: true},
					smdx.PointElement{Id: Off25, Offset: 25, Type: typelabel.Uint16, Access: "rw", Mandatory: true},
					smdx.PointElement{Id: Off26, Offset: 26, Type: typelabel.Uint16, Access: "rw", Mandatory: true},
					smdx.PointElement{Id: Off27, Offset: 27, Type: typelabel.Uint16, Access: "rw", Mandatory: true},
					smdx.PointElement{Id: Off28, Offset: 28, Type: typelabel.Uint16, Access: "rw", Mandatory: true},
					smdx.PointElement{Id: Off29, Offset: 29, Type: typelabel.Uint16, Access: "rw", Mandatory: true},
					smdx.PointElement{Id: Off30, Offset: 30, Type: typelabel.Uint16, Access: "rw", Mandatory: true},
					smdx.PointElement{Id: Off31, Offset: 31, Type: typelabel.Uint16, Access: "rw", Mandatory: true},
					smdx.PointElement{Id: Off32, Offset: 32, Type: typelabel.Uint16, Access: "rw", Mandatory: true},
					smdx.PointElement{Id: Off33, Offset: 33, Type: typelabel.Uint16, Access: "rw", Mandatory: true},
					smdx.PointElement{Id: Off34, Offset: 34, Type: typelabel.Uint16, Access: "rw", Mandatory: true},
					smdx.PointElement{Id: Off35, Offset: 35, Type: typelabel.Uint16, Access: "rw", Mandatory: true},
					smdx.PointElement{Id: Off36, Offset: 36, Type: typelabel.Uint16, Access: "rw", Mandatory: true},
					smdx.PointElement{Id: Off37, Offset: 37, Type: typelabel.Uint16, Access: "rw", Mandatory: true},
					smdx.PointElement{Id: Off38, Offset: 38, Type: typelabel.Uint16, Access: "rw", Mandatory: true},
					smdx.PointElement{Id: Off39, Offset: 39, Type: typelabel.Uint16, Access: "rw", Mandatory: true},
					smdx.PointElement{Id: Off40, Offset: 40, Type: typelabel.Uint16, Access: "rw", Mandatory: true},
					smdx.PointElement{Id: Off41, Offset: 41, Type: typelabel.Uint16, Access: "rw", Mandatory: true},
					smdx.PointElement{Id: Off42, Offset: 42, Type: typelabel.Uint16, Access: "rw", Mandatory: true},
					smdx.PointElement{Id: Off43, Offset: 43, Type: typelabel.Uint16, Access: "rw", Mandatory: true},
					smdx.PointElement{Id: Off44, Offset: 44, Type: typelabel.Uint16, Access: "rw", Mandatory: true},
					smdx.PointElement{Id: Off45, Offset: 45, Type: typelabel.Uint16, Access: "rw", Mandatory: true},
					smdx.PointElement{Id: Off46, Offset: 46, Type: typelabel.Uint16, Access: "rw", Mandatory: true},
					smdx.PointElement{Id: Off47, Offset: 47, Type: typelabel.Uint16, Access: "rw", Mandatory: true},
					smdx.PointElement{Id: Off48, Offset: 48, Type: typelabel.Uint16, Access: "rw", Mandatory: true},
					smdx.PointElement{Id: Off49, Offset: 49, Type: typelabel.Uint16, Access: "rw", Mandatory: true},
					smdx.PointElement{Id: Off50, Offset: 50, Type: typelabel.Uint16, Access: "rw", Mandatory: true},
					smdx.PointElement{Id: Ts, Offset: 51, Type: typelabel.Uint32, Access: "rw", Mandatory: true, Label: "Timestamp", Description: "Timestamp value is the number of seconds since January 1, 2000"},
					smdx.PointElement{Id: Ms, Offset: 53, Type: typelabel.Uint16, Access: "rw", Mandatory: true, Label: "Milliseconds", Description: "Millisecond counter 0-999"},
					smdx.PointElement{Id: Seq, Offset: 54, Type: typelabel.Uint16, Access: "rw", Mandatory: true, Label: "Sequence", Description: "Sequence number of request"},
					smdx.PointElement{Id: Role, Offset: 55, Type: typelabel.Uint16, Access: "rw", Mandatory: true, Label: "Role", Description: "Digital Signature ID"},
					smdx.PointElement{Id: Alg, Offset: 56, Type: typelabel.Enum16, Access: "r", Mandatory: true, Label: "Algorithm", Description: "Algorithm used to compute the digital signature"},
					smdx.PointElement{Id: N, Offset: 57, Type: typelabel.Uint16, Access: "r", Mandatory: true, Label: "N", Description: "Number of registers comprising the digital signature."},
				},
			},
			smdx.BlockElement{
				Length: 1,
				Type:   "repeating",
				Points: []smdx.PointElement{
					smdx.PointElement{Id: DS, Offset: 0, Type: typelabel.Uint16, Access: "r", Mandatory: true, Label: "DS", Description: "Digital Signature"},
				},
			},
		}})
}
