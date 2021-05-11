package model64001

// Code generated by ../../generators/models.go. DO NOT EDIT.
// You can regenerate it by running 'go generate ./models' from the directory above.

import (
	"github.com/andig/gosunspec"
	"github.com/andig/gosunspec/smdx"
	"github.com/andig/gosunspec/typelabel"
)

// Block64001 - Veris Status and Configuration -

const (
	ModelID          = 64001
	ModelLabel       = "Veris Status and Configuration"
	ModelDescription = ""
)

const (
	Boots    = "Boots"
	Cmd      = "Cmd"
	Config   = "Config"
	HWRev    = "HWRev"
	LEDblink = "LEDblink"
	LEDon    = "LEDon"
	Loc      = "Loc"
	OSFWRev  = "OSFWRev"
	ProdRev  = "ProdRev"
	RSFWRev  = "RSFWRev"
	Reserved = "Reserved"
	S1Addr   = "S1Addr"
	S1ID     = "S1ID"
	S1OSVer  = "S1OSVer"
	S1Serial = "S1Serial"
	S1Ver    = "S1Ver"
	S2Addr   = "S2Addr"
	S2ID     = "S2ID"
	S2OSVer  = "S2OSVer"
	S2Serial = "S2Serial"
	S2Ver    = "S2Ver"
	S3Addr   = "S3Addr"
	S3ID     = "S3ID"
	S3OSVer  = "S3OSVer"
	S3Serial = "S3Serial"
	S3Ver    = "S3Ver"
	S4Addr   = "S4Addr"
	S4ID     = "S4ID"
	S4OSVer  = "S4OSVer"
	S4Serial = "S4Serial"
	S4Ver    = "S4Ver"
	Sensors  = "Sensors"
	Status   = "Status"
	Switch   = "Switch"
	Talking  = "Talking"
)

type Block64001 struct {
	Cmd      sunspec.Enum16     `sunspec:"offset=0,access=rw"`
	HWRev    uint16             `sunspec:"offset=1"`
	RSFWRev  uint16             `sunspec:"offset=2"`
	OSFWRev  uint16             `sunspec:"offset=3"`
	ProdRev  string             `sunspec:"offset=4,len=2"`
	Boots    uint16             `sunspec:"offset=6"`
	Switch   sunspec.Bitfield16 `sunspec:"offset=7"`
	Sensors  uint16             `sunspec:"offset=8"`
	Talking  uint16             `sunspec:"offset=9"`
	Status   sunspec.Bitfield16 `sunspec:"offset=10"`
	Config   sunspec.Bitfield16 `sunspec:"offset=11"`
	LEDblink uint16             `sunspec:"offset=12"`
	LEDon    uint16             `sunspec:"offset=13"`
	Reserved uint16             `sunspec:"offset=14"`
	Loc      string             `sunspec:"offset=15,len=16"`
	S1ID     sunspec.Enum16     `sunspec:"offset=31"`
	S1Addr   uint16             `sunspec:"offset=32"`
	S1OSVer  uint16             `sunspec:"offset=33"`
	S1Ver    string             `sunspec:"offset=34,len=2"`
	S1Serial string             `sunspec:"offset=36,len=5"`
	S2ID     sunspec.Enum16     `sunspec:"offset=41"`
	S2Addr   uint16             `sunspec:"offset=42"`
	S2OSVer  uint16             `sunspec:"offset=43"`
	S2Ver    string             `sunspec:"offset=44,len=2"`
	S2Serial string             `sunspec:"offset=46,len=5"`
	S3ID     sunspec.Enum16     `sunspec:"offset=51"`
	S3Addr   uint16             `sunspec:"offset=52"`
	S3OSVer  uint16             `sunspec:"offset=53"`
	S3Ver    string             `sunspec:"offset=54,len=2"`
	S3Serial string             `sunspec:"offset=56,len=5"`
	S4ID     sunspec.Enum16     `sunspec:"offset=61"`
	S4Addr   uint16             `sunspec:"offset=62"`
	S4OSVer  uint16             `sunspec:"offset=63"`
	S4Ver    string             `sunspec:"offset=64,len=2"`
	S4Serial string             `sunspec:"offset=66,len=5"`
}

func (block *Block64001) GetId() sunspec.ModelId {
	return ModelID
}

func init() {
	smdx.RegisterModel(&smdx.ModelElement{
		Id:     ModelID,
		Name:   "",
		Length: 71,
		Blocks: []smdx.BlockElement{
			{
				Length: 71,
				Points: []smdx.PointElement{
					{Id: Cmd, Offset: 0, Type: typelabel.Enum16, Access: "rw", Label: "Command Code", Description: ""},
					{Id: HWRev, Offset: 1, Type: typelabel.Uint16, Label: "Hardware Revision", Description: ""},
					{Id: RSFWRev, Offset: 2, Type: typelabel.Uint16, Label: "RS FW Revision", Description: ""},
					{Id: OSFWRev, Offset: 3, Type: typelabel.Uint16, Label: "OS FW Revision", Description: ""},
					{Id: ProdRev, Offset: 4, Type: typelabel.String, Length: 2, Label: "Product Revision", Description: ""},
					{Id: Boots, Offset: 6, Type: typelabel.Uint16, Label: "Boot Count", Description: ""},
					{Id: Switch, Offset: 7, Type: typelabel.Bitfield16, Label: "DIP Switches", Description: ""},
					{Id: Sensors, Offset: 8, Type: typelabel.Uint16, Label: "Num Detected Sensors", Description: ""},
					{Id: Talking, Offset: 9, Type: typelabel.Uint16, Label: "Num Communicating Sensors", Description: ""},
					{Id: Status, Offset: 10, Type: typelabel.Bitfield16, Label: "System Status", Description: ""},
					{Id: Config, Offset: 11, Type: typelabel.Bitfield16, Label: "System Configuration", Description: ""},
					{Id: LEDblink, Offset: 12, Type: typelabel.Uint16, Units: "Pct", Label: "LED Blink Threshold", Description: ""},
					{Id: LEDon, Offset: 13, Type: typelabel.Uint16, Units: "Pct", Label: "LED On Threshold", Description: ""},
					{Id: Reserved, Offset: 14, Type: typelabel.Uint16},
					{Id: Loc, Offset: 15, Type: typelabel.String, Length: 16, Label: "Location String", Description: ""},
					{Id: S1ID, Offset: 31, Type: typelabel.Enum16, Label: "Sensor 1 Unit ID", Description: ""},
					{Id: S1Addr, Offset: 32, Type: typelabel.Uint16, Label: "Sensor 1 Address", Description: ""},
					{Id: S1OSVer, Offset: 33, Type: typelabel.Uint16, Label: "Sensor 1 OS Version", Description: ""},
					{Id: S1Ver, Offset: 34, Type: typelabel.String, Length: 2, Label: "Sensor 1 Product Version", Description: ""},
					{Id: S1Serial, Offset: 36, Type: typelabel.String, Length: 5, Label: "Sensor 1 Serial Num", Description: ""},
					{Id: S2ID, Offset: 41, Type: typelabel.Enum16, Label: "Sensor 2 Unit ID", Description: ""},
					{Id: S2Addr, Offset: 42, Type: typelabel.Uint16, Label: "Sensor 2 Address", Description: ""},
					{Id: S2OSVer, Offset: 43, Type: typelabel.Uint16, Label: "Sensor 2 OS Version", Description: ""},
					{Id: S2Ver, Offset: 44, Type: typelabel.String, Length: 2, Label: "Sensor 2 Product Version", Description: ""},
					{Id: S2Serial, Offset: 46, Type: typelabel.String, Length: 5, Label: "Sensor 2 Serial Num", Description: ""},
					{Id: S3ID, Offset: 51, Type: typelabel.Enum16, Label: "Sensor 3 Unit ID", Description: ""},
					{Id: S3Addr, Offset: 52, Type: typelabel.Uint16, Label: "Sensor 3 Address", Description: ""},
					{Id: S3OSVer, Offset: 53, Type: typelabel.Uint16, Label: "Sensor 3 OS Version", Description: ""},
					{Id: S3Ver, Offset: 54, Type: typelabel.String, Length: 2, Label: "Sensor 3 Product Version", Description: ""},
					{Id: S3Serial, Offset: 56, Type: typelabel.String, Length: 5, Label: "Sensor 3 Serial Num", Description: ""},
					{Id: S4ID, Offset: 61, Type: typelabel.Enum16, Label: "Sensor 4 Unit ID", Description: ""},
					{Id: S4Addr, Offset: 62, Type: typelabel.Uint16, Label: "Sensor 4 Address", Description: ""},
					{Id: S4OSVer, Offset: 63, Type: typelabel.Uint16, Label: "Sensor 4 OS Version", Description: ""},
					{Id: S4Ver, Offset: 64, Type: typelabel.String, Length: 2, Label: "Sensor 4 Product Version", Description: ""},
					{Id: S4Serial, Offset: 66, Type: typelabel.String, Length: 5, Label: "Sensor 4 Serial Num", Description: ""},
				},
			},
		}})
}
