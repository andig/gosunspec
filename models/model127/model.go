package model127

// Code generated by ../../generators/models.go. DO NOT EDIT.
// You can regenerate it by running 'go generate ./models' from the directory above.

import (
	"github.com/andig/gosunspec"
	"github.com/andig/gosunspec/smdx"
	"github.com/andig/gosunspec/typelabel"
)

// Block127 - Freq-Watt Param - Parameterized Frequency-Watt

const (
	ModelID          = 127
	ModelLabel       = "Freq-Watt Param"
	ModelDescription = "Parameterized Frequency-Watt "
)

const (
	HysEna       = "HysEna"
	HzStop       = "HzStop"
	HzStopWGra   = "HzStopWGra"
	HzStr        = "HzStr"
	HzStrStop_SF = "HzStrStop_SF"
	ModEna       = "ModEna"
	Pad          = "Pad"
	RmpIncDec_SF = "RmpIncDec_SF"
	WGra         = "WGra"
	WGra_SF      = "WGra_SF"
)

type Block127 struct {
	WGra         uint16              `sunspec:"offset=0,len=1,sf=WGra_SF,access=rw"`
	HzStr        int16               `sunspec:"offset=1,len=1,sf=HzStrStop_SF,access=rw"`
	HzStop       int16               `sunspec:"offset=2,len=1,sf=HzStrStop_SF,access=rw"`
	HysEna       sunspec.Bitfield16  `sunspec:"offset=3,len=1,access=rw"`
	ModEna       sunspec.Bitfield16  `sunspec:"offset=4,len=1,access=rw"`
	HzStopWGra   uint16              `sunspec:"offset=5,len=1,sf=RmpIncDec_SF,access=rw"`
	WGra_SF      sunspec.ScaleFactor `sunspec:"offset=6,len=1,access=r"`
	HzStrStop_SF sunspec.ScaleFactor `sunspec:"offset=7,len=1,access=r"`
	RmpIncDec_SF sunspec.ScaleFactor `sunspec:"offset=8,len=1,access=r"`
	Pad          sunspec.Pad         `sunspec:"offset=9,len=1,access=r"`
}

func (block *Block127) GetId() sunspec.ModelId {
	return ModelID
}

func init() {
	smdx.RegisterModel(&smdx.ModelElement{
		Id:     ModelID,
		Name:   "freq_watt_param",
		Length: 10,
		Blocks: []smdx.BlockElement{
			{
				Length: 10,
				Type:   "fixed",
				Points: []smdx.PointElement{
					{Id: WGra, Offset: 0, Type: typelabel.Uint16, ScaleFactor: "WGra_SF", Units: "% PM/Hz", Access: "rw", Length: 1, Mandatory: true, Label: "WGra", Description: "The slope of the reduction in the maximum allowed watts output as a function of frequency."},
					{Id: HzStr, Offset: 1, Type: typelabel.Int16, ScaleFactor: "HzStrStop_SF", Units: "Hz", Access: "rw", Length: 1, Mandatory: true, Label: "HzStr", Description: "The frequency deviation from nominal frequency (ECPNomHz) at which a snapshot of the instantaneous power output is taken to act as the CAPPED power level (PM) and above which reduction in power output occurs."},
					{Id: HzStop, Offset: 2, Type: typelabel.Int16, ScaleFactor: "HzStrStop_SF", Units: "Hz", Access: "rw", Length: 1, Mandatory: true, Label: "HzStop", Description: "The frequency deviation from nominal frequency (ECPNomHz) at which curtailed power output may return to normal and the cap on the power level value is removed."},
					{Id: HysEna, Offset: 3, Type: typelabel.Bitfield16, Access: "rw", Length: 1, Mandatory: true, Label: "HysEna", Description: "Enable hysteresis"},
					{Id: ModEna, Offset: 4, Type: typelabel.Bitfield16, Access: "rw", Length: 1, Mandatory: true, Label: "ModEna", Description: "Is Parameterized Frequency-Watt control active."},
					{Id: HzStopWGra, Offset: 5, Type: typelabel.Uint16, ScaleFactor: "RmpIncDec_SF", Units: "% WMax/min", Access: "rw", Length: 1, Label: "HzStopWGra", Description: "The maximum time-based rate of change at which power output returns to normal after having been capped by an over frequency event."},
					{Id: WGra_SF, Offset: 6, Type: typelabel.ScaleFactor, Access: "r", Length: 1, Label: "WGra_SF", Description: "Scale factor for output gradient."},
					{Id: HzStrStop_SF, Offset: 7, Type: typelabel.ScaleFactor, Access: "r", Length: 1, Label: "HzStrStop_SF", Description: "Scale factor for frequency deviations."},
					{Id: RmpIncDec_SF, Offset: 8, Type: typelabel.ScaleFactor, Access: "r", Length: 1, Label: "RmpIncDec_SF", Description: "Scale factor for increment and decrement ramps."},
					{Id: Pad, Offset: 9, Type: typelabel.Pad, Access: "r", Length: 1},
				},
			},
		}})
}
