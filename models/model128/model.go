package model128

// Code generated by ../../generators/models.go. DO NOT EDIT.
// You can regenerate it by running 'go generate ./models' from the directory above.

import (
	"github.com/andig/gosunspec"
	"github.com/andig/gosunspec/smdx"
	"github.com/andig/gosunspec/typelabel"
)

// Block128 - Dynamic Reactive Current - Dynamic Reactive Current

const (
	ModelID          = 128
	ModelLabel       = "Dynamic Reactive Current"
	ModelDescription = "Dynamic Reactive Current "
)

const (
	ArGraMod   = "ArGraMod"
	ArGraSag   = "ArGraSag"
	ArGraSwell = "ArGraSwell"
	ArGra_SF   = "ArGra_SF"
	BlkZnTmms  = "BlkZnTmms"
	BlkZnV     = "BlkZnV"
	DbVMax     = "DbVMax"
	DbVMin     = "DbVMin"
	FilTms     = "FilTms"
	HoldTmms   = "HoldTmms"
	HysBlkZnV  = "HysBlkZnV"
	ModEna     = "ModEna"
	Pad        = "Pad"
	VRefPct_SF = "VRefPct_SF"
)

type Block128 struct {
	ArGraMod   sunspec.Enum16      `sunspec:"offset=0,len=1,access=rw"`
	ArGraSag   uint16              `sunspec:"offset=1,len=1,sf=ArGra_SF,access=rw"`
	ArGraSwell uint16              `sunspec:"offset=2,len=1,sf=ArGra_SF,access=rw"`
	ModEna     sunspec.Bitfield16  `sunspec:"offset=3,len=1,access=rw"`
	FilTms     uint16              `sunspec:"offset=4,len=1,access=rw"`
	DbVMin     uint16              `sunspec:"offset=5,len=1,sf=VRefPct_SF,access=rw"`
	DbVMax     uint16              `sunspec:"offset=6,len=1,sf=VRefPct_SF,access=rw"`
	BlkZnV     uint16              `sunspec:"offset=7,len=1,sf=VRefPct_SF,access=rw"`
	HysBlkZnV  uint16              `sunspec:"offset=8,len=1,sf=VRefPct_SF,access=rw"`
	BlkZnTmms  uint16              `sunspec:"offset=9,len=1,access=rw"`
	HoldTmms   uint16              `sunspec:"offset=10,len=1,access=rw"`
	ArGra_SF   sunspec.ScaleFactor `sunspec:"offset=11,len=1,access=r"`
	VRefPct_SF sunspec.ScaleFactor `sunspec:"offset=12,len=1,access=r"`
	Pad        sunspec.Pad         `sunspec:"offset=13,len=1,access=r"`
}

func (block *Block128) GetId() sunspec.ModelId {
	return ModelID
}

func init() {
	smdx.RegisterModel(&smdx.ModelElement{
		Id:     ModelID,
		Name:   "reactive_current",
		Length: 14,
		Blocks: []smdx.BlockElement{
			{
				Length: 14,
				Points: []smdx.PointElement{
					{Id: ArGraMod, Offset: 0, Type: typelabel.Enum16, Access: "rw", Length: 1, Mandatory: true, Label: "ArGraMod", Description: "Indicates if gradients trend toward zero at the edges of the deadband or trend toward zero at the center of the deadband."},
					{Id: ArGraSag, Offset: 1, Type: typelabel.Uint16, ScaleFactor: "ArGra_SF", Units: "%ARtg/%dV", Access: "rw", Length: 1, Mandatory: true, Label: "ArGraSag", Description: "The gradient used to increase capacitive dynamic current. A value of 0 indicates no additional reactive current support."},
					{Id: ArGraSwell, Offset: 2, Type: typelabel.Uint16, ScaleFactor: "ArGra_SF", Units: "%ARtg/%dV", Access: "rw", Length: 1, Mandatory: true, Label: "ArGraSwell", Description: "The gradient used to increase inductive dynamic current.  A value of 0 indicates no additional reactive current support."},
					{Id: ModEna, Offset: 3, Type: typelabel.Bitfield16, Access: "rw", Length: 1, Mandatory: true, Label: "ModEna", Description: "Activate dynamic reactive current model"},
					{Id: FilTms, Offset: 4, Type: typelabel.Uint16, Units: "Secs", Access: "rw", Length: 1, Label: "FilTms", Description: "The time window used to calculate the moving average voltage."},
					{Id: DbVMin, Offset: 5, Type: typelabel.Uint16, ScaleFactor: "VRefPct_SF", Units: "% VRef", Access: "rw", Length: 1, Label: "DbVMin", Description: "The lower delta voltage limit for which negative voltage deviations less than this value no dynamic vars are produced."},
					{Id: DbVMax, Offset: 6, Type: typelabel.Uint16, ScaleFactor: "VRefPct_SF", Units: "% VRef", Access: "rw", Length: 1, Label: "DbVMax", Description: "The upper delta voltage limit for which positive voltage deviations less than this value no dynamic current produced."},
					{Id: BlkZnV, Offset: 7, Type: typelabel.Uint16, ScaleFactor: "VRefPct_SF", Units: "% VRef", Access: "rw", Length: 1, Label: "BlkZnV", Description: "Block zone voltage which defines a lower voltage boundary below which no dynamic current is produced."},
					{Id: HysBlkZnV, Offset: 8, Type: typelabel.Uint16, ScaleFactor: "VRefPct_SF", Units: "% VRef", Access: "rw", Length: 1, Label: "HysBlkZnV", Description: "Hysteresis voltage used with BlkZnV."},
					{Id: BlkZnTmms, Offset: 9, Type: typelabel.Uint16, Units: "mSecs", Access: "rw", Length: 1, Label: "BlkZnTmms", Description: "Block zone time the time before which reactive current support remains active regardless of how low the voltage drops."},
					{Id: HoldTmms, Offset: 10, Type: typelabel.Uint16, Units: "mSecs", Access: "rw", Length: 1, Label: "HoldTmms", Description: "Hold time during which reactive current support continues after the average voltage has entered the dead zone."},
					{Id: ArGra_SF, Offset: 11, Type: typelabel.ScaleFactor, Access: "r", Length: 1, Mandatory: true, Label: "ArGra_SF", Description: "Scale factor for the gradients."},
					{Id: VRefPct_SF, Offset: 12, Type: typelabel.ScaleFactor, Access: "r", Length: 1, Label: "VRefPct_SF", Description: "Scale factor for the voltage zone and limit settings."},
					{Id: Pad, Offset: 13, Type: typelabel.Pad, Access: "r", Length: 1},
				},
			},
		}})
}
