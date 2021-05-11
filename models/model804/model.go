package model804

// Code generated by ../../generators/models.go. DO NOT EDIT.
// You can regenerate it by running 'go generate ./models' from the directory above.

import (
	"github.com/andig/gosunspec"
	"github.com/andig/gosunspec/smdx"
	"github.com/andig/gosunspec/typelabel"
)

// Block804 - Lithium-Ion String Model -

const (
	ModelID          = 804
	ModelLabel       = "Lithium-Ion String Model"
	ModelDescription = ""
)

const (
	A                 = "A"
	A_SF              = "A_SF"
	CellVAvg          = "CellVAvg"
	CellVMax          = "CellVMax"
	CellVMaxMod       = "CellVMaxMod"
	CellVMin          = "CellVMin"
	CellVMinMod       = "CellVMinMod"
	CellV_SF          = "CellV_SF"
	ConFail           = "ConFail"
	ConSt             = "ConSt"
	DoD               = "DoD"
	DoD_SF            = "DoD_SF"
	Evt1              = "Evt1"
	Evt2              = "Evt2"
	EvtVnd1           = "EvtVnd1"
	EvtVnd2           = "EvtVnd2"
	Idx               = "Idx"
	ModCellTmpAvg     = "ModCellTmpAvg"
	ModCellTmpMax     = "ModCellTmpMax"
	ModCellTmpMaxCell = "ModCellTmpMaxCell"
	ModCellTmpMin     = "ModCellTmpMin"
	ModCellTmpMinCell = "ModCellTmpMinCell"
	ModCellVAvg       = "ModCellVAvg"
	ModCellVMax       = "ModCellVMax"
	ModCellVMaxCell   = "ModCellVMaxCell"
	ModCellVMin       = "ModCellVMin"
	ModCellVMinCell   = "ModCellVMinCell"
	ModNCell          = "ModNCell"
	ModSoC            = "ModSoC"
	ModSoH            = "ModSoH"
	ModTmpAvg         = "ModTmpAvg"
	ModTmpMax         = "ModTmpMax"
	ModTmpMaxMod      = "ModTmpMaxMod"
	ModTmpMin         = "ModTmpMin"
	ModTmpMinMod      = "ModTmpMinMod"
	ModTmp_SF         = "ModTmp_SF"
	NCellBal          = "NCellBal"
	NCyc              = "NCyc"
	NMod              = "NMod"
	Pad1              = "Pad1"
	Pad2              = "Pad2"
	Pad3              = "Pad3"
	Pad4              = "Pad4"
	Pad5              = "Pad5"
	Pad6              = "Pad6"
	Pad7              = "Pad7"
	SetCon            = "SetCon"
	SetEna            = "SetEna"
	SoC               = "SoC"
	SoC_SF            = "SoC_SF"
	SoH               = "SoH"
	SoH_SF            = "SoH_SF"
	St                = "St"
	V                 = "V"
	V_SF              = "V_SF"
)

type Block804Repeat struct {
	ModNCell          uint16      `sunspec:"offset=0"`
	ModSoC            uint16      `sunspec:"offset=1,sf=SoC_SF"`
	ModSoH            uint16      `sunspec:"offset=2,sf=SoH_SF"`
	ModCellVMax       uint16      `sunspec:"offset=3,sf=CellV_SF"`
	ModCellVMaxCell   uint16      `sunspec:"offset=4"`
	ModCellVMin       uint16      `sunspec:"offset=5,sf=CellV_SF"`
	ModCellVMinCell   uint16      `sunspec:"offset=6,sf=CellV_SF"`
	ModCellVAvg       uint16      `sunspec:"offset=7,sf=CellV_SF"`
	ModCellTmpMax     int16       `sunspec:"offset=8,sf=ModTmp_SF"`
	ModCellTmpMaxCell uint16      `sunspec:"offset=9"`
	ModCellTmpMin     int16       `sunspec:"offset=10,sf=ModTmp_SF"`
	ModCellTmpMinCell uint16      `sunspec:"offset=11"`
	ModCellTmpAvg     int16       `sunspec:"offset=12,sf=ModTmp_SF"`
	Pad5              sunspec.Pad `sunspec:"offset=13"`
	Pad6              sunspec.Pad `sunspec:"offset=14"`
	Pad7              sunspec.Pad `sunspec:"offset=15"`
}

type Block804 struct {
	Idx          uint16              `sunspec:"offset=0"`
	NMod         uint16              `sunspec:"offset=1"`
	St           sunspec.Bitfield32  `sunspec:"offset=2"`
	ConFail      sunspec.Enum16      `sunspec:"offset=4"`
	NCellBal     uint16              `sunspec:"offset=5"`
	SoC          uint16              `sunspec:"offset=6,sf=SoC_SF"`
	DoD          uint16              `sunspec:"offset=7,sf=DoD_SF"`
	NCyc         uint32              `sunspec:"offset=8"`
	SoH          uint16              `sunspec:"offset=10,sf=SoH_SF"`
	A            int16               `sunspec:"offset=11,sf=A_SF"`
	V            uint16              `sunspec:"offset=12,sf=V_SF"`
	CellVMax     uint16              `sunspec:"offset=13,sf=CellV_SF"`
	CellVMaxMod  uint16              `sunspec:"offset=14"`
	CellVMin     uint16              `sunspec:"offset=15,sf=CellV_SF"`
	CellVMinMod  uint16              `sunspec:"offset=16"`
	CellVAvg     uint16              `sunspec:"offset=17,sf=CellV_SF"`
	ModTmpMax    int16               `sunspec:"offset=18,sf=ModTmp_SF"`
	ModTmpMaxMod uint16              `sunspec:"offset=19"`
	ModTmpMin    int16               `sunspec:"offset=20,sf=ModTmp_SF"`
	ModTmpMinMod uint16              `sunspec:"offset=21"`
	ModTmpAvg    int16               `sunspec:"offset=22,sf=ModTmp_SF"`
	Pad1         sunspec.Pad         `sunspec:"offset=23"`
	ConSt        sunspec.Bitfield32  `sunspec:"offset=24"`
	Evt1         sunspec.Bitfield32  `sunspec:"offset=26"`
	Evt2         sunspec.Bitfield32  `sunspec:"offset=28"`
	EvtVnd1      sunspec.Bitfield32  `sunspec:"offset=30"`
	EvtVnd2      sunspec.Bitfield32  `sunspec:"offset=32"`
	SetEna       sunspec.Enum16      `sunspec:"offset=34,access=rw"`
	SetCon       sunspec.Enum16      `sunspec:"offset=35,access=rw"`
	SoC_SF       sunspec.ScaleFactor `sunspec:"offset=36"`
	SoH_SF       sunspec.ScaleFactor `sunspec:"offset=37"`
	DoD_SF       sunspec.ScaleFactor `sunspec:"offset=38"`
	A_SF         sunspec.ScaleFactor `sunspec:"offset=39"`
	V_SF         sunspec.ScaleFactor `sunspec:"offset=40"`
	CellV_SF     sunspec.ScaleFactor `sunspec:"offset=41"`
	ModTmp_SF    sunspec.ScaleFactor `sunspec:"offset=42"`
	Pad2         sunspec.Pad         `sunspec:"offset=43"`
	Pad3         sunspec.Pad         `sunspec:"offset=44"`
	Pad4         sunspec.Pad         `sunspec:"offset=45"`

	Repeats []Block804Repeat
}

func (block *Block804) GetId() sunspec.ModelId {
	return ModelID
}

func init() {
	smdx.RegisterModel(&smdx.ModelElement{
		Id:     ModelID,
		Name:   "lithium_ion_string",
		Length: 62,
		Blocks: []smdx.BlockElement{
			{
				Length: 46,
				Points: []smdx.PointElement{
					{Id: Idx, Offset: 0, Type: typelabel.Uint16, Mandatory: true, Label: "String Index", Description: "Index of the string within the bank."},
					{Id: NMod, Offset: 1, Type: typelabel.Uint16, Mandatory: true, Label: "Module Count", Description: "Count of modules in the string."},
					{Id: St, Offset: 2, Type: typelabel.Bitfield32, Mandatory: true, Label: "String Status", Description: "Current status of the string."},
					{Id: ConFail, Offset: 4, Type: typelabel.Enum16, Label: "Connection Failure Reason", Description: ""},
					{Id: NCellBal, Offset: 5, Type: typelabel.Uint16, Label: "String Cell Balancing Count", Description: "Number of cells currently being balanced in the string."},
					{Id: SoC, Offset: 6, Type: typelabel.Uint16, ScaleFactor: "SoC_SF", Units: "%", Mandatory: true, Label: "String State of Charge", Description: "Battery string state of charge, expressed as a percentage."},
					{Id: DoD, Offset: 7, Type: typelabel.Uint16, ScaleFactor: "DoD_SF", Units: "%", Label: "String Depth of Discharge", Description: "Depth of discharge for the string, expressed as a percentage."},
					{Id: NCyc, Offset: 8, Type: typelabel.Uint32, Label: "String Cycle Count", Description: "Number of discharge cycles executed upon the string."},
					{Id: SoH, Offset: 10, Type: typelabel.Uint16, ScaleFactor: "SoH_SF", Units: "%", Label: "String State of Health", Description: "Battery string state of health, expressed as a percentage."},
					{Id: A, Offset: 11, Type: typelabel.Int16, ScaleFactor: "A_SF", Units: "A", Mandatory: true, Label: "String Current", Description: "String current measurement."},
					{Id: V, Offset: 12, Type: typelabel.Uint16, ScaleFactor: "V_SF", Units: "V", Label: "String Voltage", Description: "String voltage measurement."},
					{Id: CellVMax, Offset: 13, Type: typelabel.Uint16, ScaleFactor: "CellV_SF", Units: "V", Mandatory: true, Label: "Max Cell Voltage", Description: "Maximum voltage for all cells in the string."},
					{Id: CellVMaxMod, Offset: 14, Type: typelabel.Uint16, Label: "Max Cell Voltage Module", Description: "Module containing the cell with maximum cell voltage."},
					{Id: CellVMin, Offset: 15, Type: typelabel.Uint16, ScaleFactor: "CellV_SF", Units: "V", Mandatory: true, Label: "Min Cell Voltage", Description: "Minimum voltage for all cells in the string."},
					{Id: CellVMinMod, Offset: 16, Type: typelabel.Uint16, Label: "Min Cell Voltage Module", Description: "Module containing the cell with minimum cell voltage."},
					{Id: CellVAvg, Offset: 17, Type: typelabel.Uint16, ScaleFactor: "CellV_SF", Units: "V", Mandatory: true, Label: "Average Cell Voltage", Description: "Average voltage for all cells in the string."},
					{Id: ModTmpMax, Offset: 18, Type: typelabel.Int16, ScaleFactor: "ModTmp_SF", Units: "C", Mandatory: true, Label: "Max Module Temperature", Description: "Maximum temperature for all modules in the string."},
					{Id: ModTmpMaxMod, Offset: 19, Type: typelabel.Uint16, Mandatory: true, Label: "Max Module Temperature Module", Description: "Module with the maximum temperature."},
					{Id: ModTmpMin, Offset: 20, Type: typelabel.Int16, ScaleFactor: "ModTmp_SF", Units: "C", Mandatory: true, Label: "Min Module Temperature", Description: "Minimum temperature for all modules in the string."},
					{Id: ModTmpMinMod, Offset: 21, Type: typelabel.Uint16, Mandatory: true, Label: "Min Module Temperature Module", Description: "Module with the minimum temperature."},
					{Id: ModTmpAvg, Offset: 22, Type: typelabel.Int16, ScaleFactor: "ModTmp_SF", Units: "C", Mandatory: true, Label: "Average Module Temperature", Description: "Average temperature for all modules in the string."},
					{Id: Pad1, Offset: 23, Type: typelabel.Pad, Mandatory: true, Label: "Pad", Description: "Pad register."},
					{Id: ConSt, Offset: 24, Type: typelabel.Bitfield32, Label: "Contactor Status", Description: "Status of the contactor(s) for the string."},
					{Id: Evt1, Offset: 26, Type: typelabel.Bitfield32, Mandatory: true, Label: "String Event 1", Description: "Alarms, warnings and status values.  Bit flags."},
					{Id: Evt2, Offset: 28, Type: typelabel.Bitfield32, Label: "String Event 2", Description: "Alarms, warnings and status values.  Bit flags."},
					{Id: EvtVnd1, Offset: 30, Type: typelabel.Bitfield32, Label: "Vendor Event Bitfield 1", Description: "Vendor defined events."},
					{Id: EvtVnd2, Offset: 32, Type: typelabel.Bitfield32, Label: "Vendor Event Bitfield 2", Description: "Vendor defined events."},
					{Id: SetEna, Offset: 34, Type: typelabel.Enum16, Access: "rw", Label: "Enable/Disable String", Description: "Enables and disables the string.  Should reset to 0 upon completion."},
					{Id: SetCon, Offset: 35, Type: typelabel.Enum16, Access: "rw", Label: "Connect/Disconnect String", Description: "Connects and disconnects the string."},
					{Id: SoC_SF, Offset: 36, Type: typelabel.ScaleFactor, Mandatory: true},
					{Id: SoH_SF, Offset: 37, Type: typelabel.ScaleFactor},
					{Id: DoD_SF, Offset: 38, Type: typelabel.ScaleFactor},
					{Id: A_SF, Offset: 39, Type: typelabel.ScaleFactor, Mandatory: true},
					{Id: V_SF, Offset: 40, Type: typelabel.ScaleFactor},
					{Id: CellV_SF, Offset: 41, Type: typelabel.ScaleFactor, Mandatory: true},
					{Id: ModTmp_SF, Offset: 42, Type: typelabel.ScaleFactor, Mandatory: true},
					{Id: Pad2, Offset: 43, Type: typelabel.Pad, Mandatory: true, Label: "Pad", Description: "Pad register."},
					{Id: Pad3, Offset: 44, Type: typelabel.Pad, Mandatory: true, Label: "Pad", Description: "Pad register."},
					{Id: Pad4, Offset: 45, Type: typelabel.Pad, Mandatory: true, Label: "Pad", Description: "Pad register."},
				},
			},
			{Name: "lithium_ion_string_module",
				Length: 16,
				Type:   "repeating",
				Points: []smdx.PointElement{
					{Id: ModNCell, Offset: 0, Type: typelabel.Uint16, Mandatory: true, Label: "Module Cell Count", Description: "Count of all cells in the module."},
					{Id: ModSoC, Offset: 1, Type: typelabel.Uint16, ScaleFactor: "SoC_SF", Units: "%", Label: "Module SoC", Description: "Module state of charge, expressed as a percentage."},
					{Id: ModSoH, Offset: 2, Type: typelabel.Uint16, ScaleFactor: "SoH_SF", Units: "%", Label: "Module SoH", Description: "Module state of health, expressed as a percentage."},
					{Id: ModCellVMax, Offset: 3, Type: typelabel.Uint16, ScaleFactor: "CellV_SF", Units: "V", Mandatory: true, Label: "Max Cell Voltage", Description: "Maximum voltage for all cells in the module."},
					{Id: ModCellVMaxCell, Offset: 4, Type: typelabel.Uint16, Label: "Max Cell Voltage Cell", Description: "Cell with maximum voltage."},
					{Id: ModCellVMin, Offset: 5, Type: typelabel.Uint16, ScaleFactor: "CellV_SF", Units: "V", Mandatory: true, Label: "Min Cell Voltage", Description: "Minimum voltage for all cells in the module."},
					{Id: ModCellVMinCell, Offset: 6, Type: typelabel.Uint16, ScaleFactor: "CellV_SF", Units: "V", Label: "Min Cell Voltage Cell", Description: "Cell with minimum voltage."},
					{Id: ModCellVAvg, Offset: 7, Type: typelabel.Uint16, ScaleFactor: "CellV_SF", Units: "V", Mandatory: true, Label: "Average Cell Voltage", Description: "Average voltage for all cells in the module."},
					{Id: ModCellTmpMax, Offset: 8, Type: typelabel.Int16, ScaleFactor: "ModTmp_SF", Units: "C", Mandatory: true, Label: "Max Cell Temperature", Description: "Maximum temperature for all cells in the module."},
					{Id: ModCellTmpMaxCell, Offset: 9, Type: typelabel.Uint16, Label: "Max Cell Temperature Cell", Description: "Cell with maximum temperature."},
					{Id: ModCellTmpMin, Offset: 10, Type: typelabel.Int16, ScaleFactor: "ModTmp_SF", Units: "C", Mandatory: true, Label: "Min Cell Temperature", Description: "Minimum temperature for all cells in the module."},
					{Id: ModCellTmpMinCell, Offset: 11, Type: typelabel.Uint16, Label: "Min Cell Temperature Cell", Description: "Cell with minimum temperature."},
					{Id: ModCellTmpAvg, Offset: 12, Type: typelabel.Int16, ScaleFactor: "ModTmp_SF", Units: "C", Mandatory: true, Label: "Average Cell Temperature", Description: "Average temperature for all cells in the module."},
					{Id: Pad5, Offset: 13, Type: typelabel.Pad, Mandatory: true, Label: "Pad", Description: "Pad register."},
					{Id: Pad6, Offset: 14, Type: typelabel.Pad, Mandatory: true, Label: "Pad", Description: "Pad register."},
					{Id: Pad7, Offset: 15, Type: typelabel.Pad, Mandatory: true, Label: "Pad", Description: "Pad register."},
				},
			},
		}})
}
