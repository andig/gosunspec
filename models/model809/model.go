// NOTICE
// This file was automatically generated by ../../generators/models.go. Do not edit it!
// You can regenerate it by running 'go generate ./models' from the directory above.

package model809

import (
	"github.com/andig/gosunspec"
	"github.com/andig/gosunspec/smdx"
	"github.com/andig/gosunspec/typelabel"
)

// Block809 - Flow Battery Stack Model -

const (
	ModelID          = 809
	ModelLabel       = "Flow Battery Stack Model"
	ModelDescription = ""
)

const (
	CellTBD  = "CellTBD"
	StackTBD = "StackTBD"
)

type Block809Repeat struct {
	CellTBD uint16 `sunspec:"offset=0"`
}

type Block809 struct {
	StackTBD uint16 `sunspec:"offset=0"`

	Repeats []Block809Repeat
}

func (self *Block809) GetId() sunspec.ModelId {
	return ModelID
}

func init() {
	smdx.RegisterModel(&smdx.ModelElement{
		Id:     ModelID,
		Name:   "flow_battery_stack",
		Length: 2,
		Blocks: []smdx.BlockElement{
			smdx.BlockElement{
				Length: 1,

				Points: []smdx.PointElement{
					smdx.PointElement{Id: StackTBD, Offset: 0, Type: typelabel.Uint16, Mandatory: true, Label: "Stack Points To Be Determined", Description: ""},
				},
			},
			smdx.BlockElement{Name: "cell",
				Length: 1,
				Type:   "repeating",
				Points: []smdx.PointElement{
					smdx.PointElement{Id: CellTBD, Offset: 0, Type: typelabel.Uint16, Mandatory: true, Label: "Cell Points To Be Determined", Description: ""},
				},
			},
		}})
}
