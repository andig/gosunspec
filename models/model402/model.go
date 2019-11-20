// NOTICE
// This file was automatically generated by ../../generators/models.go. Do not edit it!
// You can regenerate it by running 'go generate ./models' from the directory above.

package model402

import (
	"github.com/andig/gosunspec"
	"github.com/andig/gosunspec/smdx"
	"github.com/andig/gosunspec/typelabel"
)

// Block402 - String Combiner (Advanced) - An advanced string combiner

const (
	ModelID          = 402
	ModelLabel       = "String Combiner (Advanced)"
	ModelDescription = "An advanced string combiner"
)

const (
	DCA      = "DCA"
	DCAMax   = "DCAMax"
	DCA_SF   = "DCA_SF"
	DCAhr    = "DCAhr"
	DCAhr_SF = "DCAhr_SF"
	DCPR     = "DCPR"
	DCV      = "DCV"
	DCV_SF   = "DCV_SF"
	DCW      = "DCW"
	DCW_SF   = "DCW_SF"
	DCWh     = "DCWh"
	DCWh_SF  = "DCWh_SF"
	Evt      = "Evt"
	EvtVnd   = "EvtVnd"
	InDCA    = "InDCA"
	InDCAhr  = "InDCAhr"
	InDCPR   = "InDCPR"
	InDCV    = "InDCV"
	InDCW    = "InDCW"
	InDCWh   = "InDCWh"
	InEvt    = "InEvt"
	InID     = "InID"
	InN      = "InN"
	N        = "N"
	Tmp      = "Tmp"
)

type Block402Repeat struct {
	InID    uint16             `sunspec:"offset=0"`
	InEvt   sunspec.Bitfield32 `sunspec:"offset=1"`
	EvtVnd  sunspec.Bitfield32 `sunspec:"offset=3"`
	InDCA   int16              `sunspec:"offset=5,sf=DCA_SF"`
	InDCAhr uint32             `sunspec:"offset=6,sf=DCAhr_SF"`
	InDCV   uint16             `sunspec:"offset=8,sf=DCV_SF"`
	InDCW   int16              `sunspec:"offset=9,sf=DCWh_SF"`
	InDCWh  uint32             `sunspec:"offset=10"`
	InDCPR  uint16             `sunspec:"offset=12"`
	InN     uint16             `sunspec:"offset=13"`
}

type Block402 struct {
	DCA_SF   sunspec.ScaleFactor `sunspec:"offset=0"`
	DCAhr_SF sunspec.ScaleFactor `sunspec:"offset=1"`
	DCV_SF   sunspec.ScaleFactor `sunspec:"offset=2"`
	DCW_SF   sunspec.ScaleFactor `sunspec:"offset=3"`
	DCWh_SF  sunspec.ScaleFactor `sunspec:"offset=4"`
	DCAMax   uint16              `sunspec:"offset=5"`
	N        sunspec.Count       `sunspec:"offset=6"`
	Evt      sunspec.Bitfield32  `sunspec:"offset=7"`
	EvtVnd   sunspec.Bitfield32  `sunspec:"offset=9"`
	DCA      int16               `sunspec:"offset=11,sf=DCA_SF"`
	DCAhr    uint32              `sunspec:"offset=12,sf=DCAhr_SF"`
	DCV      uint16              `sunspec:"offset=14,sf=DCV_SF"`
	Tmp      int16               `sunspec:"offset=15"`
	DCW      int16               `sunspec:"offset=16,sf=DCW_SF"`
	DCPR     uint16              `sunspec:"offset=17"`
	DCWh     uint32              `sunspec:"offset=18,sf=DCWh_SF"`

	Repeats []Block402Repeat
}

func (self *Block402) GetId() sunspec.ModelId {
	return ModelID
}

func init() {
	smdx.RegisterModel(&smdx.ModelElement{
		Id:     ModelID,
		Name:   "string_combiner",
		Length: 34,
		Blocks: []smdx.BlockElement{
			smdx.BlockElement{
				Length: 20,

				Points: []smdx.PointElement{
					smdx.PointElement{Id: DCA_SF, Offset: 0, Type: typelabel.ScaleFactor, Mandatory: true},
					smdx.PointElement{Id: DCAhr_SF, Offset: 1, Type: typelabel.ScaleFactor},
					smdx.PointElement{Id: DCV_SF, Offset: 2, Type: typelabel.ScaleFactor},
					smdx.PointElement{Id: DCW_SF, Offset: 3, Type: typelabel.ScaleFactor},
					smdx.PointElement{Id: DCWh_SF, Offset: 4, Type: typelabel.ScaleFactor, Mandatory: true},
					smdx.PointElement{Id: DCAMax, Offset: 5, Type: typelabel.Uint16, Units: "A", Label: "Rating", Description: "Maximum DC Current Rating"},
					smdx.PointElement{Id: N, Offset: 6, Type: typelabel.Count, Label: "N", Description: "Number of Inputs"},
					smdx.PointElement{Id: Evt, Offset: 7, Type: typelabel.Bitfield32, Mandatory: true, Label: "Event", Description: "Bitmask value.  Events"},
					smdx.PointElement{Id: EvtVnd, Offset: 9, Type: typelabel.Bitfield32, Label: "Vendor Event", Description: "Bitmask value.  Vendor defined events"},
					smdx.PointElement{Id: DCA, Offset: 11, Type: typelabel.Int16, ScaleFactor: "DCA_SF", Units: "A", Mandatory: true, Label: "Amps", Description: "Total measured current"},
					smdx.PointElement{Id: DCAhr, Offset: 12, Type: typelabel.Uint32, ScaleFactor: "DCAhr_SF", Units: "Ah", Label: "Amp-hours", Description: "Total metered Amp-hours"},
					smdx.PointElement{Id: DCV, Offset: 14, Type: typelabel.Uint16, ScaleFactor: "DCV_SF", Units: "V", Label: "Voltage", Description: "Output Voltage"},
					smdx.PointElement{Id: Tmp, Offset: 15, Type: typelabel.Int16, Units: "C", Label: "Temp", Description: "Internal operating temperature"},
					smdx.PointElement{Id: DCW, Offset: 16, Type: typelabel.Int16, ScaleFactor: "DCW_SF", Units: "W", Label: "Watts", Description: "Output power"},
					smdx.PointElement{Id: DCPR, Offset: 17, Type: typelabel.Uint16, Units: "Pct", Label: "PR", Description: "DC Performance ratio value"},
					smdx.PointElement{Id: DCWh, Offset: 18, Type: typelabel.Uint32, ScaleFactor: "DCWh_SF", Units: "Wh", Mandatory: true, Label: "Watt-hours", Description: "Output energy"},
				},
			},
			smdx.BlockElement{Name: "string",
				Length: 14,
				Type:   "repeating",
				Points: []smdx.PointElement{
					smdx.PointElement{Id: InID, Offset: 0, Type: typelabel.Uint16, Mandatory: true, Label: "ID", Description: "Uniquely identifies this input set"},
					smdx.PointElement{Id: InEvt, Offset: 1, Type: typelabel.Bitfield32, Mandatory: true, Label: "Input Event", Description: "String Input Event Flags"},
					smdx.PointElement{Id: EvtVnd, Offset: 3, Type: typelabel.Bitfield32, Label: "Vendor Event", Description: "Bitmask value.  Vendor defined events"},
					smdx.PointElement{Id: InDCA, Offset: 5, Type: typelabel.Int16, ScaleFactor: "DCA_SF", Units: "A", Mandatory: true, Label: "Amps", Description: "String Input Current"},
					smdx.PointElement{Id: InDCAhr, Offset: 6, Type: typelabel.Uint32, ScaleFactor: "DCAhr_SF", Units: "Ah", Label: "Amp-hours", Description: "String Input Amp-Hours"},
					smdx.PointElement{Id: InDCV, Offset: 8, Type: typelabel.Uint16, ScaleFactor: "DCV_SF", Units: "V", Label: "Voltage", Description: "String Input Voltage"},
					smdx.PointElement{Id: InDCW, Offset: 9, Type: typelabel.Int16, ScaleFactor: "DCWh_SF", Units: "W", Label: "Watts", Description: "String Input Power"},
					smdx.PointElement{Id: InDCWh, Offset: 10, Type: typelabel.Uint32, Units: "Wh", Label: "Watt-hours", Description: "String Input Energy"},
					smdx.PointElement{Id: InDCPR, Offset: 12, Type: typelabel.Uint16, Units: "Pct", Label: "PR", Description: "String Performance Ratio"},
					smdx.PointElement{Id: InN, Offset: 13, Type: typelabel.Uint16, Label: "N", Description: "Number of modules in this input string"},
				},
			},
		}})
}
