// NOTICE
// This file was automatically generated by ../../generators/models.go. Do not edit it!
// You can regenerate it by running 'go generate ./models' from the directory above.

package model501

import (
	"github.com/andig/gosunspec"
	"github.com/andig/gosunspec/smdx"
	"github.com/andig/gosunspec/typelabel"
)

// Block501 - Solar Module - A solar module model supporting DC-DC converter

const (
	ModelID          = 501
	ModelLabel       = "Solar Module"
	ModelDescription = "A solar module model supporting DC-DC converter"
)

const (
	Ctl      = "Ctl"
	CtlVal   = "CtlVal"
	CtlVend  = "CtlVend"
	Evt      = "Evt"
	EvtVend  = "EvtVend"
	InA      = "InA"
	InV      = "InV"
	InW      = "InW"
	InWh     = "InWh"
	OutA     = "OutA"
	OutV     = "OutV"
	OutW     = "OutW"
	OutWh    = "OutWh"
	Stat     = "Stat"
	StatVend = "StatVend"
	Tmp      = "Tmp"
	Tms      = "Tms"
)

type Block501 struct {
	Stat     sunspec.Enum16     `sunspec:"offset=0"`
	StatVend sunspec.Enum16     `sunspec:"offset=1"`
	Evt      sunspec.Bitfield32 `sunspec:"offset=2"`
	EvtVend  sunspec.Bitfield32 `sunspec:"offset=4"`
	Ctl      sunspec.Enum16     `sunspec:"offset=6,access=rw"`
	CtlVend  sunspec.Enum32     `sunspec:"offset=7,access=rw"`
	CtlVal   int32              `sunspec:"offset=9,access=rw"`
	Tms      uint32             `sunspec:"offset=11"`
	OutA     float32            `sunspec:"offset=13"`
	OutV     float32            `sunspec:"offset=15"`
	OutWh    float32            `sunspec:"offset=17"`
	OutW     float32            `sunspec:"offset=19"`
	Tmp      float32            `sunspec:"offset=21"`
	InA      float32            `sunspec:"offset=23"`
	InV      float32            `sunspec:"offset=25"`
	InWh     float32            `sunspec:"offset=27"`
	InW      float32            `sunspec:"offset=29"`
}

func (self *Block501) GetId() sunspec.ModelId {
	return ModelID
}

func init() {
	smdx.RegisterModel(&smdx.ModelElement{
		Id:     ModelID,
		Name:   "solar_module",
		Length: 31,
		Blocks: []smdx.BlockElement{
			smdx.BlockElement{
				Length: 31,

				Points: []smdx.PointElement{
					smdx.PointElement{Id: Stat, Offset: 0, Type: typelabel.Enum16, Mandatory: true, Label: "Status", Description: "Enumerated value.  Module Status Code"},
					smdx.PointElement{Id: StatVend, Offset: 1, Type: typelabel.Enum16, Label: "Vendor Status", Description: "Module Vendor Status Code"},
					smdx.PointElement{Id: Evt, Offset: 2, Type: typelabel.Bitfield32, Mandatory: true, Label: "Events", Description: "Bitmask value.  Module Event Flags"},
					smdx.PointElement{Id: EvtVend, Offset: 4, Type: typelabel.Bitfield32, Label: "Vendor Module Event Flags", Description: "Vendor specific flags"},
					smdx.PointElement{Id: Ctl, Offset: 6, Type: typelabel.Enum16, Access: "rw", Label: "Control", Description: "Module Control"},
					smdx.PointElement{Id: CtlVend, Offset: 7, Type: typelabel.Enum32, Access: "rw", Label: "Vendor Control", Description: "Vendor Module Control"},
					smdx.PointElement{Id: CtlVal, Offset: 9, Type: typelabel.Int32, Access: "rw", Label: "Control Value", Description: "Module Control Value"},
					smdx.PointElement{Id: Tms, Offset: 11, Type: typelabel.Uint32, Units: "Secs", Label: "Timestamp", Description: "Time in seconds since 2000 epoch"},
					smdx.PointElement{Id: OutA, Offset: 13, Type: typelabel.Float32, Units: "A", Label: "Output Current", Description: "Output Current"},
					smdx.PointElement{Id: OutV, Offset: 15, Type: typelabel.Float32, Units: "V", Label: "Output Voltage", Description: "Output Voltage"},
					smdx.PointElement{Id: OutWh, Offset: 17, Type: typelabel.Float32, Units: "Wh", Label: "Output Energy", Description: "Output Energy"},
					smdx.PointElement{Id: OutW, Offset: 19, Type: typelabel.Float32, Units: "W", Label: "Output Power", Description: "Output Power"},
					smdx.PointElement{Id: Tmp, Offset: 21, Type: typelabel.Float32, Units: "C", Label: "Temp", Description: "Module Temperature"},
					smdx.PointElement{Id: InA, Offset: 23, Type: typelabel.Float32, Units: "A", Label: "Input Current", Description: "Input Current"},
					smdx.PointElement{Id: InV, Offset: 25, Type: typelabel.Float32, Units: "V", Label: "Input Voltage", Description: "Input Voltage"},
					smdx.PointElement{Id: InWh, Offset: 27, Type: typelabel.Float32, Units: "Wh", Label: "Input Energy", Description: "Input Energy"},
					smdx.PointElement{Id: InW, Offset: 29, Type: typelabel.Float32, Units: "W", Label: "Input Power", Description: "Input Power"},
				},
			},
		}})
}
