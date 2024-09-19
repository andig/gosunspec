package model705

// Code generated by ../../generators/models.go. DO NOT EDIT.
// You can regenerate it by running 'go generate ./models' from the directory above.

import (
	"github.com/andig/gosunspec"
	"github.com/andig/gosunspec/smdx"
	"github.com/andig/gosunspec/typelabel"
)

// Block705 - DER Volt-Var - DER Volt-Var model.

const (
	ModelID          = 705
	ModelLabel       = "DER Volt-Var"
	ModelDescription = "DER Volt-Var model."
)

const (
	ActPt       = "ActPt"
	AdptCrvReq  = "AdptCrvReq"
	AdptCrvRslt = "AdptCrvRslt"
	DeptRef     = "DeptRef"
	DeptRef_SF  = "DeptRef_SF"
	Ena         = "Ena"
	NCrv        = "NCrv"
	NPt         = "NPt"
	Pri         = "Pri"
	ReadOnly    = "ReadOnly"
	RspTms      = "RspTms"
	RspTms_SF   = "RspTms_SF"
	RvrtCrv     = "RvrtCrv"
	RvrtRem     = "RvrtRem"
	RvrtTms     = "RvrtTms"
	VRef        = "VRef"
	VRefAuto    = "VRefAuto"
	VRefAutoEna = "VRefAutoEna"
	VRefAutoTms = "VRefAutoTms"
	V_SF        = "V_SF"
)

type Block705Repeat struct {
	ActPt       uint16         `sunspec:"offset=0,len=1,access=rw"`
	DeptRef     sunspec.Enum16 `sunspec:"offset=1,len=1,access=rw"`
	Pri         sunspec.Enum16 `sunspec:"offset=2,len=1,access=rw"`
	VRef        uint16         `sunspec:"offset=3,len=1,access=rw"`
	VRefAuto    uint16         `sunspec:"offset=4,len=1,access=r"`
	VRefAutoEna sunspec.Enum16 `sunspec:"offset=5,len=1,access=rw"`
	VRefAutoTms uint16         `sunspec:"offset=6,len=1,access=rw"`
	RspTms      uint32         `sunspec:"offset=7,len=2,sf=RspTms_SF,access=rw"`
	ReadOnly    sunspec.Enum16 `sunspec:"offset=9,len=1,access=r"`
}

type Block705 struct {
	Ena         sunspec.Enum16      `sunspec:"offset=0,len=1,access=rw"`
	AdptCrvReq  uint16              `sunspec:"offset=1,len=1,access=rw"`
	AdptCrvRslt sunspec.Enum16      `sunspec:"offset=2,len=1,access=r"`
	NPt         uint16              `sunspec:"offset=3,len=1,access=r"`
	NCrv        uint16              `sunspec:"offset=4,len=1,access=r"`
	RvrtTms     uint32              `sunspec:"offset=5,len=2,access=rw"`
	RvrtRem     uint32              `sunspec:"offset=7,len=2,access=r"`
	RvrtCrv     uint16              `sunspec:"offset=9,len=1,access=rw"`
	V_SF        sunspec.ScaleFactor `sunspec:"offset=10,len=1,access=r"`
	DeptRef_SF  sunspec.ScaleFactor `sunspec:"offset=11,len=1,access=r"`
	RspTms_SF   sunspec.ScaleFactor `sunspec:"offset=12,len=1,access=r"`

	Repeats []Block705Repeat
}

func (block *Block705) GetId() sunspec.ModelId {
	return ModelID
}

func init() {
	smdx.RegisterModel(&smdx.ModelElement{
		Id:     ModelID,
		Name:   "DERVoltVar",
		Length: 23,
		Blocks: []smdx.BlockElement{
			{
				Length: 13,
				Points: []smdx.PointElement{
					{Id: Ena, Offset: 0, Type: typelabel.Enum16, Access: "rw", Length: 1, Mandatory: true, Label: "Module Enable", Description: "Volt-Var control enable."},
					{Id: AdptCrvReq, Offset: 1, Type: typelabel.Uint16, Access: "rw", Length: 1, Mandatory: true, Label: "Adopt Curve Request", Description: "Index of  curve points to adopt. First curve index is 1."},
					{Id: AdptCrvRslt, Offset: 2, Type: typelabel.Enum16, Access: "r", Length: 1, Mandatory: true, Label: "Adopt Curve Result", Description: "Result of last adopt curve operation."},
					{Id: NPt, Offset: 3, Type: typelabel.Uint16, Access: "r", Length: 1, Mandatory: true, Label: "Number Of Points", Description: "Number of curve points supported."},
					{Id: NCrv, Offset: 4, Type: typelabel.Uint16, Access: "r", Length: 1, Mandatory: true, Label: "Stored Curve Count", Description: "Number of stored curves supported."},
					{Id: RvrtTms, Offset: 5, Type: typelabel.Uint32, Access: "rw", Length: 2, Label: "Reversion Timeout", Description: "Reversion time in seconds.  0 = No reversion time."},
					{Id: RvrtRem, Offset: 7, Type: typelabel.Uint32, Access: "r", Length: 2, Label: "Reversion Time Remaining", Description: "Reversion time remaining in seconds."},
					{Id: RvrtCrv, Offset: 9, Type: typelabel.Uint16, Access: "rw", Length: 1, Label: "Reversion Curve", Description: "Default curve after reversion timeout."},
					{Id: V_SF, Offset: 10, Type: typelabel.ScaleFactor, Access: "r", Length: 1, Mandatory: true, Label: "Voltage Scale Factor", Description: "Scale factor for curve voltage points."},
					{Id: DeptRef_SF, Offset: 11, Type: typelabel.ScaleFactor, Access: "r", Length: 1, Mandatory: true, Label: "Var Scale Factor", Description: "Scale factor for curve var points."},
					{Id: RspTms_SF, Offset: 12, Type: typelabel.ScaleFactor, Access: "r", Length: 1, Mandatory: true, Label: "Open-Loop Scale Factor", Description: "Open loop response time scale factor."},
				},
			},
			{Name: "Crv",
				Length: 10,
				Type:   "repeating",
				Points: []smdx.PointElement{
					{Id: ActPt, Offset: 0, Type: typelabel.Uint16, Access: "rw", Length: 1, Mandatory: true, Label: "Active Points", Description: "Number of active points."},
					{Id: DeptRef, Offset: 1, Type: typelabel.Enum16, Access: "rw", Length: 1, Mandatory: true, Label: "Dependent Reference", Description: "Curve dependent reference."},
					{Id: Pri, Offset: 2, Type: typelabel.Enum16, Access: "rw", Length: 1},
					{Id: VRef, Offset: 3, Type: typelabel.Uint16, Access: "rw", Length: 1, Label: "Vref Adjustment", Description: "Vref adjustment as percentage of nominal voltage."},
					{Id: VRefAuto, Offset: 4, Type: typelabel.Uint16, Access: "r", Length: 1, Label: "Current Autonomous Vref", Description: "Autonomous vref value as a percentage of nominal voltage."},
					{Id: VRefAutoEna, Offset: 5, Type: typelabel.Enum16, Access: "rw", Length: 1, Label: "Autonomous Vref Enable", Description: "Enable automonous vref."},
					{Id: VRefAutoTms, Offset: 6, Type: typelabel.Uint16, Access: "rw", Length: 1, Label: "Auto Vref Time Constant", Description: "Autonomous vref time constant."},
					{Id: RspTms, Offset: 7, Type: typelabel.Uint32, ScaleFactor: "RspTms_SF", Units: "Secs", Access: "rw", Length: 2, Label: "Open Loop Response Time", Description: "Open loop response time."},
					{Id: ReadOnly, Offset: 9, Type: typelabel.Enum16, Access: "r", Length: 1, Label: "Curve Access", Description: "Curve read-write access."},
				},
			},
		}})
}