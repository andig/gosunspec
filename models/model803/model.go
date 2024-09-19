package model803

// Code generated by ../../generators/models.go. DO NOT EDIT.
// You can regenerate it by running 'go generate ./models' from the directory above.

import (
	"github.com/andig/gosunspec"
	"github.com/andig/gosunspec/smdx"
	"github.com/andig/gosunspec/typelabel"
)

// Block803 - Lithium-Ion Battery Bank Model -

const (
	ModelID          = 803
	ModelLabel       = "Lithium-Ion Battery Bank Model"
	ModelDescription = ""
)

const (
	A_SF            = "A_SF"
	CellV_SF        = "CellV_SF"
	ModTmpAvg       = "ModTmpAvg"
	ModTmpMax       = "ModTmpMax"
	ModTmpMaxMod    = "ModTmpMaxMod"
	ModTmpMaxStr    = "ModTmpMaxStr"
	ModTmpMin       = "ModTmpMin"
	ModTmpMinMod    = "ModTmpMinMod"
	ModTmpMinStr    = "ModTmpMinStr"
	ModTmp_SF       = "ModTmp_SF"
	NCellBal        = "NCellBal"
	NStr            = "NStr"
	NStrCon         = "NStrCon"
	Pad1            = "Pad1"
	Pad2            = "Pad2"
	SoC_SF          = "SoC_SF"
	SoH_SF          = "SoH_SF"
	StrA            = "StrA"
	StrAAvg         = "StrAAvg"
	StrAMax         = "StrAMax"
	StrAMaxStr      = "StrAMaxStr"
	StrAMin         = "StrAMin"
	StrAMinStr      = "StrAMinStr"
	StrCellVAvg     = "StrCellVAvg"
	StrCellVMax     = "StrCellVMax"
	StrCellVMaxMod  = "StrCellVMaxMod"
	StrCellVMin     = "StrCellVMin"
	StrCellVMinMod  = "StrCellVMinMod"
	StrConFail      = "StrConFail"
	StrConSt        = "StrConSt"
	StrDisRsn       = "StrDisRsn"
	StrEvt1         = "StrEvt1"
	StrEvt2         = "StrEvt2"
	StrEvtVnd1      = "StrEvtVnd1"
	StrEvtVnd2      = "StrEvtVnd2"
	StrModTmpAvg    = "StrModTmpAvg"
	StrModTmpMax    = "StrModTmpMax"
	StrModTmpMaxMod = "StrModTmpMaxMod"
	StrModTmpMin    = "StrModTmpMin"
	StrModTmpMinMod = "StrModTmpMinMod"
	StrNMod         = "StrNMod"
	StrSetCon       = "StrSetCon"
	StrSetEna       = "StrSetEna"
	StrSoC          = "StrSoC"
	StrSoH          = "StrSoH"
	StrSt           = "StrSt"
	StrVAvg         = "StrVAvg"
	StrVMax         = "StrVMax"
	StrVMaxStr      = "StrVMaxStr"
	StrVMin         = "StrVMin"
	StrVMinStr      = "StrVMinStr"
	V_SF            = "V_SF"
)

type Block803Repeat struct {
	StrNMod         uint16             `sunspec:"offset=0,len=1,access=r"`
	StrSt           sunspec.Bitfield32 `sunspec:"offset=1,len=2,access=r"`
	StrConFail      sunspec.Enum16     `sunspec:"offset=3,len=1,access=r"`
	StrSoC          uint16             `sunspec:"offset=4,len=1,sf=SoC_SF,access=r"`
	StrSoH          uint16             `sunspec:"offset=5,len=1,sf=SoH_SF,access=r"`
	StrA            int16              `sunspec:"offset=6,len=1,sf=A_SF,access=r"`
	StrCellVMax     uint16             `sunspec:"offset=7,len=1,sf=CellV_SF,access=r"`
	StrCellVMaxMod  uint16             `sunspec:"offset=8,len=1,access=r"`
	StrCellVMin     uint16             `sunspec:"offset=9,len=1,sf=CellV_SF,access=r"`
	StrCellVMinMod  uint16             `sunspec:"offset=10,len=1,access=r"`
	StrCellVAvg     uint16             `sunspec:"offset=11,len=1,sf=CellV_SF,access=r"`
	StrModTmpMax    int16              `sunspec:"offset=12,len=1,sf=ModTmp_SF,access=r"`
	StrModTmpMaxMod uint16             `sunspec:"offset=13,len=1,access=r"`
	StrModTmpMin    int16              `sunspec:"offset=14,len=1,sf=ModTmp_SF,access=r"`
	StrModTmpMinMod uint16             `sunspec:"offset=15,len=1,access=r"`
	StrModTmpAvg    int16              `sunspec:"offset=16,len=1,sf=ModTmp_SF,access=r"`
	StrDisRsn       sunspec.Enum16     `sunspec:"offset=17,len=1,access=r"`
	StrConSt        sunspec.Bitfield32 `sunspec:"offset=18,len=2,access=r"`
	StrEvt1         sunspec.Bitfield32 `sunspec:"offset=20,len=2,access=r"`
	StrEvt2         sunspec.Bitfield32 `sunspec:"offset=22,len=2,access=r"`
	StrEvtVnd1      sunspec.Bitfield32 `sunspec:"offset=24,len=2,access=r"`
	StrEvtVnd2      sunspec.Bitfield32 `sunspec:"offset=26,len=2,access=r"`
	StrSetEna       sunspec.Enum16     `sunspec:"offset=28,len=1,access=rw"`
	StrSetCon       sunspec.Enum16     `sunspec:"offset=29,len=1,access=rw"`
	Pad1            sunspec.Pad        `sunspec:"offset=30,len=1,access=r"`
	Pad2            sunspec.Pad        `sunspec:"offset=31,len=1,access=r"`
}

type Block803 struct {
	NStr         uint16              `sunspec:"offset=0,len=1,access=r"`
	NStrCon      uint16              `sunspec:"offset=1,len=1,access=r"`
	ModTmpMax    int16               `sunspec:"offset=2,len=1,sf=ModTmp_SF,access=r"`
	ModTmpMaxStr uint16              `sunspec:"offset=3,len=1,access=r"`
	ModTmpMaxMod uint16              `sunspec:"offset=4,len=1,access=r"`
	ModTmpMin    int16               `sunspec:"offset=5,len=1,sf=ModTmp_SF,access=r"`
	ModTmpMinStr uint16              `sunspec:"offset=6,len=1,access=r"`
	ModTmpMinMod uint16              `sunspec:"offset=7,len=1,access=r"`
	ModTmpAvg    uint16              `sunspec:"offset=8,len=1,access=r"`
	StrVMax      uint16              `sunspec:"offset=9,len=1,sf=V_SF,access=r"`
	StrVMaxStr   uint16              `sunspec:"offset=10,len=1,access=r"`
	StrVMin      uint16              `sunspec:"offset=11,len=1,sf=V_SF,access=r"`
	StrVMinStr   uint16              `sunspec:"offset=12,len=1,access=r"`
	StrVAvg      uint16              `sunspec:"offset=13,len=1,sf=V_SF,access=r"`
	StrAMax      int16               `sunspec:"offset=14,len=1,sf=A_SF,access=r"`
	StrAMaxStr   uint16              `sunspec:"offset=15,len=1,access=r"`
	StrAMin      int16               `sunspec:"offset=16,len=1,sf=A_SF,access=r"`
	StrAMinStr   uint16              `sunspec:"offset=17,len=1,access=r"`
	StrAAvg      int16               `sunspec:"offset=18,len=1,sf=A_SF,access=r"`
	NCellBal     uint16              `sunspec:"offset=19,len=1,access=r"`
	CellV_SF     sunspec.ScaleFactor `sunspec:"offset=20,len=1,access=r"`
	ModTmp_SF    sunspec.ScaleFactor `sunspec:"offset=21,len=1,access=r"`
	A_SF         sunspec.ScaleFactor `sunspec:"offset=22,len=1,access=r"`
	SoH_SF       sunspec.ScaleFactor `sunspec:"offset=23,len=1,access=r"`
	SoC_SF       sunspec.ScaleFactor `sunspec:"offset=24,len=1,access=r"`
	V_SF         sunspec.ScaleFactor `sunspec:"offset=25,len=1,access=r"`

	Repeats []Block803Repeat
}

func (block *Block803) GetId() sunspec.ModelId {
	return ModelID
}

func init() {
	smdx.RegisterModel(&smdx.ModelElement{
		Id:     ModelID,
		Name:   "lithium_ion_bank",
		Length: 58,
		Blocks: []smdx.BlockElement{
			{
				Length: 26,
				Points: []smdx.PointElement{
					{Id: NStr, Offset: 0, Type: typelabel.Uint16, Access: "r", Length: 1, Mandatory: true, Label: "String Count", Description: "Number of strings in the bank."},
					{Id: NStrCon, Offset: 1, Type: typelabel.Uint16, Access: "r", Length: 1, Mandatory: true, Label: "Connected String Count", Description: "Number of strings with contactor closed."},
					{Id: ModTmpMax, Offset: 2, Type: typelabel.Int16, ScaleFactor: "ModTmp_SF", Access: "r", Length: 1, Mandatory: true, Label: "Max Module Temperature", Description: "Maximum temperature for all modules in the bank."},
					{Id: ModTmpMaxStr, Offset: 3, Type: typelabel.Uint16, Access: "r", Length: 1, Label: "Max Module Temperature String", Description: "String containing the module with maximum temperature."},
					{Id: ModTmpMaxMod, Offset: 4, Type: typelabel.Uint16, Access: "r", Length: 1, Label: "Max Module Temperature Module", Description: "Module with maximum temperature."},
					{Id: ModTmpMin, Offset: 5, Type: typelabel.Int16, ScaleFactor: "ModTmp_SF", Units: "C", Access: "r", Length: 1, Mandatory: true, Label: "Min Module Temperature", Description: "Minimum temperature for all modules in the bank."},
					{Id: ModTmpMinStr, Offset: 6, Type: typelabel.Uint16, Access: "r", Length: 1, Label: "Min Module Temperature String", Description: "String containing the module with minimum temperature."},
					{Id: ModTmpMinMod, Offset: 7, Type: typelabel.Uint16, Access: "r", Length: 1, Label: "Min Module Temperature Module", Description: "Module with minimum temperature."},
					{Id: ModTmpAvg, Offset: 8, Type: typelabel.Uint16, Access: "r", Length: 1, Label: "Average Module Temperature", Description: "Average temperature for all modules in the bank."},
					{Id: StrVMax, Offset: 9, Type: typelabel.Uint16, ScaleFactor: "V_SF", Units: "V", Access: "r", Length: 1, Label: "Max String Voltage", Description: "Maximum string voltage for all strings in the bank."},
					{Id: StrVMaxStr, Offset: 10, Type: typelabel.Uint16, Access: "r", Length: 1, Label: "Max String Voltage String", Description: "String with maximum voltage."},
					{Id: StrVMin, Offset: 11, Type: typelabel.Uint16, ScaleFactor: "V_SF", Units: "V", Access: "r", Length: 1, Label: "Min String Voltage", Description: "Minimum string voltage for all strings in the bank."},
					{Id: StrVMinStr, Offset: 12, Type: typelabel.Uint16, Access: "r", Length: 1, Label: "Min String Voltage String", Description: "String with minimum voltage."},
					{Id: StrVAvg, Offset: 13, Type: typelabel.Uint16, ScaleFactor: "V_SF", Units: "V", Access: "r", Length: 1, Label: "Average String Voltage", Description: "Average string voltage for all strings in the bank."},
					{Id: StrAMax, Offset: 14, Type: typelabel.Int16, ScaleFactor: "A_SF", Units: "A", Access: "r", Length: 1, Label: "Max String Current", Description: "Maximum current of any string in the bank."},
					{Id: StrAMaxStr, Offset: 15, Type: typelabel.Uint16, Access: "r", Length: 1, Label: "Max String Current String", Description: "String with the maximum current."},
					{Id: StrAMin, Offset: 16, Type: typelabel.Int16, ScaleFactor: "A_SF", Units: "A", Access: "r", Length: 1, Label: "Min String Current", Description: "Minimum current of any string in the bank."},
					{Id: StrAMinStr, Offset: 17, Type: typelabel.Uint16, Access: "r", Length: 1, Label: "Min String Current String", Description: "String with the minimum current."},
					{Id: StrAAvg, Offset: 18, Type: typelabel.Int16, ScaleFactor: "A_SF", Units: "A", Access: "r", Length: 1, Label: "Average String Current", Description: "Average string current for all strings in the bank."},
					{Id: NCellBal, Offset: 19, Type: typelabel.Uint16, Access: "r", Length: 1, Label: "Battery Cell Balancing Count", Description: "Total number of cells that are currently being balanced."},
					{Id: CellV_SF, Offset: 20, Type: typelabel.ScaleFactor, Access: "r", Length: 1, Mandatory: true},
					{Id: ModTmp_SF, Offset: 21, Type: typelabel.ScaleFactor, Access: "r", Length: 1, Mandatory: true},
					{Id: A_SF, Offset: 22, Type: typelabel.ScaleFactor, Access: "r", Length: 1, Mandatory: true},
					{Id: SoH_SF, Offset: 23, Type: typelabel.ScaleFactor, Access: "r", Length: 1},
					{Id: SoC_SF, Offset: 24, Type: typelabel.ScaleFactor, Access: "r", Length: 1, Mandatory: true},
					{Id: V_SF, Offset: 25, Type: typelabel.ScaleFactor, Access: "r", Length: 1},
				},
			},
			{Name: "string",
				Length: 32,
				Type:   "repeating",
				Points: []smdx.PointElement{
					{Id: StrNMod, Offset: 0, Type: typelabel.Uint16, Access: "r", Length: 1, Mandatory: true, Label: "Module Count", Description: "Count of modules in the string."},
					{Id: StrSt, Offset: 1, Type: typelabel.Bitfield32, Access: "r", Length: 2, Mandatory: true, Label: "String Status", Description: "Current status of the string."},
					{Id: StrConFail, Offset: 3, Type: typelabel.Enum16, Access: "r", Length: 1, Label: "Connection Failure Reason", Description: ""},
					{Id: StrSoC, Offset: 4, Type: typelabel.Uint16, ScaleFactor: "SoC_SF", Units: "%", Access: "r", Length: 1, Mandatory: true, Label: "String State of Charge", Description: "Battery string state of charge, expressed as a percentage."},
					{Id: StrSoH, Offset: 5, Type: typelabel.Uint16, ScaleFactor: "SoH_SF", Units: "%", Access: "r", Length: 1, Label: "String State of Health", Description: "Battery string state of health, expressed as a percentage."},
					{Id: StrA, Offset: 6, Type: typelabel.Int16, ScaleFactor: "A_SF", Units: "A", Access: "r", Length: 1, Mandatory: true, Label: "String Current", Description: "String current measurement."},
					{Id: StrCellVMax, Offset: 7, Type: typelabel.Uint16, ScaleFactor: "CellV_SF", Units: "V", Access: "r", Length: 1, Mandatory: true, Label: "Max Cell Voltage", Description: "Maximum voltage for all cells in the string."},
					{Id: StrCellVMaxMod, Offset: 8, Type: typelabel.Uint16, Access: "r", Length: 1, Label: "Max Cell Voltage Module", Description: "Module containing the maximum cell voltage."},
					{Id: StrCellVMin, Offset: 9, Type: typelabel.Uint16, ScaleFactor: "CellV_SF", Units: "V", Access: "r", Length: 1, Mandatory: true, Label: "Min Cell Voltage", Description: "Minimum voltage for all cells in the string."},
					{Id: StrCellVMinMod, Offset: 10, Type: typelabel.Uint16, Access: "r", Length: 1, Label: "Min Cell Voltage Module", Description: "Module containing the minimum cell voltage."},
					{Id: StrCellVAvg, Offset: 11, Type: typelabel.Uint16, ScaleFactor: "CellV_SF", Units: "V", Access: "r", Length: 1, Mandatory: true, Label: "Average Cell Voltage", Description: "Average voltage for all cells in the string."},
					{Id: StrModTmpMax, Offset: 12, Type: typelabel.Int16, ScaleFactor: "ModTmp_SF", Units: "C", Access: "r", Length: 1, Mandatory: true, Label: "Max Module Temperature", Description: "Maximum temperature for all modules in the bank."},
					{Id: StrModTmpMaxMod, Offset: 13, Type: typelabel.Uint16, Access: "r", Length: 1, Label: "Max Module Temperature Module", Description: "Module with the maximum temperature."},
					{Id: StrModTmpMin, Offset: 14, Type: typelabel.Int16, ScaleFactor: "ModTmp_SF", Units: "C", Access: "r", Length: 1, Mandatory: true, Label: "Min Module Temperature", Description: "Minimum temperature for all modules in the bank."},
					{Id: StrModTmpMinMod, Offset: 15, Type: typelabel.Uint16, Access: "r", Length: 1, Label: "Min Module Temperature Module", Description: "Module with the minimum temperature."},
					{Id: StrModTmpAvg, Offset: 16, Type: typelabel.Int16, ScaleFactor: "ModTmp_SF", Units: "C", Access: "r", Length: 1, Mandatory: true, Label: "Average Module Temperature", Description: "Average temperature for all modules in the bank."},
					{Id: StrDisRsn, Offset: 17, Type: typelabel.Enum16, Access: "r", Length: 1, Label: "Disabled Reason", Description: "Reason why the string is currently disabled."},
					{Id: StrConSt, Offset: 18, Type: typelabel.Bitfield32, Access: "r", Length: 2, Label: "Contactor Status", Description: "Status of the contactor(s) for the string."},
					{Id: StrEvt1, Offset: 20, Type: typelabel.Bitfield32, Access: "r", Length: 2, Mandatory: true, Label: "String Event 1", Description: "Alarms, warnings and status values.  Bit flags."},
					{Id: StrEvt2, Offset: 22, Type: typelabel.Bitfield32, Access: "r", Length: 2, Label: "String Event 2", Description: "Alarms, warnings and status values.  Bit flags."},
					{Id: StrEvtVnd1, Offset: 24, Type: typelabel.Bitfield32, Access: "r", Length: 2, Label: "Vendor String Event Bitfield 1", Description: "Vendor defined events."},
					{Id: StrEvtVnd2, Offset: 26, Type: typelabel.Bitfield32, Access: "r", Length: 2, Label: "Vendor String Event Bitfield 2", Description: "Vendor defined events."},
					{Id: StrSetEna, Offset: 28, Type: typelabel.Enum16, Access: "rw", Length: 1, Label: "Enable/Disable String", Description: "Enables and disables the string."},
					{Id: StrSetCon, Offset: 29, Type: typelabel.Enum16, Access: "rw", Length: 1, Label: "Connect/Disconnect String", Description: "Connects and disconnects the string."},
					{Id: Pad1, Offset: 30, Type: typelabel.Pad, Access: "r", Length: 1, Mandatory: true, Label: "Pad", Description: "Pad register."},
					{Id: Pad2, Offset: 31, Type: typelabel.Pad, Access: "r", Length: 1, Mandatory: true, Label: "Pad", Description: "Pad register."},
				},
			},
		}})
}
