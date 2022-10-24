package model707

// Code generated by ../../generators/models.go. DO NOT EDIT.
// You can regenerate it by running 'go generate ./models' from the directory above.

import (
	"github.com/andig/gosunspec"
	"github.com/andig/gosunspec/smdx"
	"github.com/andig/gosunspec/typelabel"
)

// Block707 - DER Trip LV - DER Low voltage trip model.

const (
	ModelID          = 707
	ModelLabel       = "DER Trip LV"
	ModelDescription = "DER Low voltage trip model."
)

const (
	AdptCrvReq  = "AdptCrvReq"
	AdptCrvRslt = "AdptCrvRslt"
	Ena         = "Ena"
	NCrvSet     = "NCrvSet"
	NPt         = "NPt"
	ReadOnly    = "ReadOnly"
	Tms_SF      = "Tms_SF"
	V_SF        = "V_SF"
)

type Block707Repeat struct {
	ReadOnly sunspec.Enum16 `sunspec:"offset=0,len=1,access=r"`
}

type Block707 struct {
	Ena         sunspec.Enum16      `sunspec:"offset=0,len=1,access=rw"`
	AdptCrvReq  uint16              `sunspec:"offset=1,len=1,access=rw"`
	AdptCrvRslt sunspec.Enum16      `sunspec:"offset=2,len=1,access=r"`
	NPt         uint16              `sunspec:"offset=3,len=1,access=r"`
	NCrvSet     uint16              `sunspec:"offset=4,len=1,access=r"`
	V_SF        sunspec.ScaleFactor `sunspec:"offset=5,len=1,access=r"`
	Tms_SF      sunspec.ScaleFactor `sunspec:"offset=6,len=1,access=r"`

	Repeats []Block707Repeat
}

func (block *Block707) GetId() sunspec.ModelId {
	return ModelID
}

func init() {
	smdx.RegisterModel(&smdx.ModelElement{
		Id:     ModelID,
		Name:   "DERTripLV",
		Length: 8,
		Blocks: []smdx.BlockElement{
			{
				Length: 7,
				Points: []smdx.PointElement{
					{Id: Ena, Offset: 0, Type: typelabel.Enum16, Access: "rw", Length: 1, Mandatory: true, Label: "DER Trip LV Module Enable", Description: "DER low voltage trip  control enable."},
					{Id: AdptCrvReq, Offset: 1, Type: typelabel.Uint16, Access: "rw", Length: 1, Mandatory: true, Label: "Adopt Curve Request", Description: "Index of  curve points to adopt. First curve index is 1."},
					{Id: AdptCrvRslt, Offset: 2, Type: typelabel.Enum16, Access: "r", Length: 1, Mandatory: true, Label: "Adopt Curve Result", Description: "Result of last adopt curve operation."},
					{Id: NPt, Offset: 3, Type: typelabel.Uint16, Access: "r", Length: 1, Mandatory: true, Label: "Number Of Points", Description: "Number of curve points supported."},
					{Id: NCrvSet, Offset: 4, Type: typelabel.Uint16, Access: "r", Length: 1, Mandatory: true, Label: "Stored Curve Count", Description: "Number of stored curves supported."},
					{Id: V_SF, Offset: 5, Type: typelabel.ScaleFactor, Access: "r", Length: 1, Mandatory: true, Label: "Voltage Scale Factor", Description: "Scale factor for curve voltage points."},
					{Id: Tms_SF, Offset: 6, Type: typelabel.ScaleFactor, Access: "r", Length: 1, Mandatory: true, Label: "Time Point Scale Factor", Description: "Scale factor for curve time points."},
				},
			},
			{Name: "Crv",
				Length: 1,
				Type:   "repeating",
				Points: []smdx.PointElement{
					{Id: ReadOnly, Offset: 0, Type: typelabel.Enum16, Access: "r", Length: 1, Label: "Curve Access", Description: "Curve read-write access."},
				},
			},
		}})
}
