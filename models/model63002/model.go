package model63002

// Code generated by ../../generators/models.go. DO NOT EDIT.
// You can regenerate it by running 'go generate ./models' from the directory above.

import (
	"github.com/andig/gosunspec"
	"github.com/andig/gosunspec/smdx"
	"github.com/andig/gosunspec/typelabel"
)

// Block63002 - SunSpec Test Model 2 -

const (
	ModelID          = 63002
	ModelLabel       = "SunSpec Test Model 2"
	ModelDescription = ""
)

const (
	Int16_1  = "int16_1"
	Int16_2  = "int16_2"
	Sunssf_1 = "sunssf_1"
	Sunssf_2 = "sunssf_2"
)

type Block63002Repeat struct {
	sunssf_1 sunspec.ScaleFactor `sunspec:"offset=0"`
	int16_1  int16               `sunspec:"offset=1,sf=sunssf_1,access=rw"`
	int16_2  int16               `sunspec:"offset=2,sf=sunssf_2"`
	sunssf_2 sunspec.ScaleFactor `sunspec:"offset=3"`
}

type Block63002 struct {
	Repeats []Block63002Repeat
}

func (block *Block63002) GetId() sunspec.ModelId {
	return ModelID
}

func init() {
	smdx.RegisterModel(&smdx.ModelElement{
		Id:     ModelID,
		Name:   "",
		Length: 4,
		Blocks: []smdx.BlockElement{
			{
				Length: 4,
				Type:   "repeating",
				Points: []smdx.PointElement{
					{Id: Sunssf_1, Offset: 0, Type: typelabel.ScaleFactor},
					{Id: Int16_1, Offset: 1, Type: typelabel.Int16, ScaleFactor: "sunssf_1", Access: "rw"},
					{Id: Int16_2, Offset: 2, Type: typelabel.Int16, ScaleFactor: "sunssf_2"},
					{Id: Sunssf_2, Offset: 3, Type: typelabel.ScaleFactor},
				},
			},
		}})
}
