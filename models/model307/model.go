package model307

// Code generated by ../../generators/models.go. DO NOT EDIT.
// You can regenerate it by running 'go generate ./models' from the directory above.

import (
	"github.com/andig/gosunspec"
	"github.com/andig/gosunspec/smdx"
	"github.com/andig/gosunspec/typelabel"
)

// Block307 - Base Met - Base Meteorological Model

const (
	ModelID          = 307
	ModelLabel       = "Base Met"
	ModelDescription = "Base Meteorological Model"
)

const (
	ElecFld = "ElecFld"
	PPT     = "PPT"
	Pres    = "Pres"
	RH      = "RH"
	Rain    = "Rain"
	Snw     = "Snw"
	SoilWet = "SoilWet"
	SurWet  = "SurWet"
	TmpAmb  = "TmpAmb"
	WndDir  = "WndDir"
	WndSpd  = "WndSpd"
)

type Block307 struct {
	TmpAmb  int16 `sunspec:"offset=0,sf=-1"`
	RH      int16 `sunspec:"offset=1"`
	Pres    int16 `sunspec:"offset=2"`
	WndSpd  int16 `sunspec:"offset=3"`
	WndDir  int16 `sunspec:"offset=4"`
	Rain    int16 `sunspec:"offset=5"`
	Snw     int16 `sunspec:"offset=6"`
	PPT     int16 `sunspec:"offset=7"`
	ElecFld int16 `sunspec:"offset=8"`
	SurWet  int16 `sunspec:"offset=9"`
	SoilWet int16 `sunspec:"offset=10"`
}

func (block *Block307) GetId() sunspec.ModelId {
	return ModelID
}

func init() {
	smdx.RegisterModel(&smdx.ModelElement{
		Id:     ModelID,
		Name:   "base_met",
		Length: 11,
		Blocks: []smdx.BlockElement{
			{
				Length: 11,
				Points: []smdx.PointElement{
					{Id: TmpAmb, Offset: 0, Type: typelabel.Int16, ScaleFactor: "-1", Units: "C", Label: "Ambient Temperature", Description: ""},
					{Id: RH, Offset: 1, Type: typelabel.Int16, Units: "Pct", Label: "Relative Humidity", Description: ""},
					{Id: Pres, Offset: 2, Type: typelabel.Int16, Units: "HPa", Label: "Barometric Pressure", Description: ""},
					{Id: WndSpd, Offset: 3, Type: typelabel.Int16, Units: "mps", Label: "Wind Speed", Description: ""},
					{Id: WndDir, Offset: 4, Type: typelabel.Int16, Units: "deg", Label: "Wind Direction", Description: ""},
					{Id: Rain, Offset: 5, Type: typelabel.Int16, Units: "mm", Label: "Rainfall", Description: ""},
					{Id: Snw, Offset: 6, Type: typelabel.Int16, Units: "mm", Label: "Snow Depth", Description: ""},
					{Id: PPT, Offset: 7, Type: typelabel.Int16, Label: "Precipitation Type", Description: " Precipitation Type (WMO 4680 SYNOP code reference)"},
					{Id: ElecFld, Offset: 8, Type: typelabel.Int16, Units: "Vm", Label: "Electric Field", Description: ""},
					{Id: SurWet, Offset: 9, Type: typelabel.Int16, Units: "kO", Label: "Surface Wetness", Description: ""},
					{Id: SoilWet, Offset: 10, Type: typelabel.Int16, Units: "Pct", Label: "Soil Wetness", Description: ""},
				},
			},
		}})
}
