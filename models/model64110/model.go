// NOTICE
// This file was automatically generated by ../../generators/models.go. Do not edit it!
// You can regenerate it by running 'go generate ./models' from the directory above.

package model64110

import (
	"github.com/andig/gosunspec"
	"github.com/andig/gosunspec/smdx"
	"github.com/andig/gosunspec/typelabel"
)

// Block64110 - OutBack AXS device -

const (
	ModelID          = 64110
	ModelLabel       = "OutBack AXS device"
	ModelDescription = ""
)

const (
	AXS_Error         = "AXS_Error"
	AXS_Spare         = "AXS_Spare"
	AXS_Status        = "AXS_Status"
	Alarm_email_addr1 = "Alarm_email_addr1"
	Alarm_email_addr2 = "Alarm_email_addr2"
	Alarm_email_en    = "Alarm_email_en"
	Alarm_email_sub   = "Alarm_email_sub"
	Ambient_temp      = "Ambient_temp"
	Battery_temp      = "Battery_temp"
	DNS1_address      = "DNS1_address"
	DNS2_address      = "DNS2_address"
	Date_Day          = "Date_Day"
	Date_month        = "Date_month"
	Date_year         = "Date_year"
	EnableDHCP        = "EnableDHCP"
	EncrypKey         = "EncrypKey"
	FTP_password      = "FTP_password"
	Gateway_address   = "Gateway_address"
	Log_mode          = "Log_mode"
	Log_retain        = "Log_retain"
	Log_write_int     = "Log_write_int"
	MAC_Address       = "MAC_Address"
	MajorFWRev        = "MajorFWRev"
	MidFWRev          = "MidFWRev"
	MinorFWRev        = "MinorFWRev"
	Modbus_port       = "Modbus_port"
	NTP_enable        = "NTP_enable"
	NTP_server_nm     = "NTP_server_nm"
	SMTP_account_nm   = "SMTP_account_nm"
	SMTP_enable_SSL   = "SMTP_enable_SSL"
	SMTP_password     = "SMTP_password"
	SMTP_server_nm    = "SMTP_server_nm"
	SMTP_user_nm      = "SMTP_user_nm"
	Stat_email_addr1  = "Stat_email_addr1"
	Stat_email_addr2  = "Stat_email_addr2"
	Stat_email_int    = "Stat_email_int"
	Stat_email_sub    = "Stat_email_sub"
	Stat_start_HR     = "Stat_start_HR"
	TCPIP_Netmask     = "TCPIP_Netmask"
	TCPIP_address     = "TCPIP_address"
	TELNET_password   = "TELNET_password"
	Temp_SF           = "Temp_SF"
	TimeZone          = "TimeZone"
	Time_hour         = "Time_hour"
	Time_minute       = "Time_minute"
	Time_second       = "Time_second"
	WritePassword     = "WritePassword"
)

type Block64110 struct {
	MajorFWRev        uint16              `sunspec:"offset=0"`
	MidFWRev          uint16              `sunspec:"offset=1"`
	MinorFWRev        uint16              `sunspec:"offset=2"`
	EncrypKey         uint16              `sunspec:"offset=3"`
	MAC_Address       string              `sunspec:"offset=4,len=7"`
	WritePassword     string              `sunspec:"offset=11,len=8"`
	EnableDHCP        sunspec.Enum16      `sunspec:"offset=19"`
	TCPIP_address     sunspec.Ipaddr      `sunspec:"offset=20"`
	Gateway_address   sunspec.Ipaddr      `sunspec:"offset=22"`
	TCPIP_Netmask     sunspec.Ipaddr      `sunspec:"offset=24"`
	DNS1_address      sunspec.Ipaddr      `sunspec:"offset=26"`
	DNS2_address      sunspec.Ipaddr      `sunspec:"offset=28"`
	Modbus_port       uint16              `sunspec:"offset=30"`
	SMTP_server_nm    string              `sunspec:"offset=31,len=20"`
	SMTP_account_nm   string              `sunspec:"offset=51,len=16"`
	SMTP_enable_SSL   sunspec.Enum16      `sunspec:"offset=67"`
	SMTP_password     string              `sunspec:"offset=68,len=8"`
	SMTP_user_nm      string              `sunspec:"offset=76,len=20"`
	Stat_email_int    uint16              `sunspec:"offset=96"`
	Stat_start_HR     uint16              `sunspec:"offset=97"`
	Stat_email_sub    string              `sunspec:"offset=98,len=25"`
	Stat_email_addr1  string              `sunspec:"offset=123,len=20"`
	Stat_email_addr2  string              `sunspec:"offset=143,len=20"`
	Alarm_email_en    sunspec.Enum16      `sunspec:"offset=163"`
	Alarm_email_sub   string              `sunspec:"offset=164,len=25"`
	Alarm_email_addr1 string              `sunspec:"offset=189,len=20"`
	Alarm_email_addr2 string              `sunspec:"offset=209,len=20"`
	FTP_password      string              `sunspec:"offset=229,len=8"`
	TELNET_password   string              `sunspec:"offset=237,len=8"`
	Log_write_int     uint16              `sunspec:"offset=245"`
	Log_retain        uint16              `sunspec:"offset=246"`
	Log_mode          sunspec.Enum16      `sunspec:"offset=247"`
	NTP_server_nm     string              `sunspec:"offset=248,len=20"`
	NTP_enable        sunspec.Enum16      `sunspec:"offset=268"`
	TimeZone          int16               `sunspec:"offset=269"`
	Date_year         uint16              `sunspec:"offset=270"`
	Date_month        uint16              `sunspec:"offset=271"`
	Date_Day          uint16              `sunspec:"offset=272"`
	Time_hour         uint16              `sunspec:"offset=273"`
	Time_minute       uint16              `sunspec:"offset=274"`
	Time_second       uint16              `sunspec:"offset=275"`
	Battery_temp      int16               `sunspec:"offset=276,sf=Temp_SF"`
	Ambient_temp      int16               `sunspec:"offset=277,sf=Temp_SF"`
	Temp_SF           sunspec.ScaleFactor `sunspec:"offset=278"`
	AXS_Error         sunspec.Bitfield16  `sunspec:"offset=279"`
	AXS_Status        sunspec.Bitfield16  `sunspec:"offset=280"`
	AXS_Spare         uint16              `sunspec:"offset=281"`
}

func (self *Block64110) GetId() sunspec.ModelId {
	return ModelID
}

func init() {
	smdx.RegisterModel(&smdx.ModelElement{
		Id:     ModelID,
		Name:   "",
		Length: 282,
		Blocks: []smdx.BlockElement{
			smdx.BlockElement{
				Length: 282,

				Points: []smdx.PointElement{
					smdx.PointElement{Id: MajorFWRev, Offset: 0, Type: typelabel.Uint16, Mandatory: true, Label: "AXS Major Firmware Number", Description: ""},
					smdx.PointElement{Id: MidFWRev, Offset: 1, Type: typelabel.Uint16, Mandatory: true, Label: "AXS Mid Firmware Number", Description: ""},
					smdx.PointElement{Id: MinorFWRev, Offset: 2, Type: typelabel.Uint16, Mandatory: true, Label: "AXS Minor Firmware Number", Description: ""},
					smdx.PointElement{Id: EncrypKey, Offset: 3, Type: typelabel.Uint16, Mandatory: true, Label: "Encryption Key", Description: ""},
					smdx.PointElement{Id: MAC_Address, Offset: 4, Type: typelabel.String, Length: 7, Mandatory: true, Label: "MAC Address", Description: ""},
					smdx.PointElement{Id: WritePassword, Offset: 11, Type: typelabel.String, Length: 8, Mandatory: true, Label: "Write Password", Description: ""},
					smdx.PointElement{Id: EnableDHCP, Offset: 19, Type: typelabel.Enum16, Mandatory: true, Label: "Enable DHCP", Description: ""},
					smdx.PointElement{Id: TCPIP_address, Offset: 20, Type: typelabel.Ipaddr, Mandatory: true, Label: "TCPIP Address", Description: ""},
					smdx.PointElement{Id: Gateway_address, Offset: 22, Type: typelabel.Ipaddr, Mandatory: true, Label: "TCPIP Gateway", Description: ""},
					smdx.PointElement{Id: TCPIP_Netmask, Offset: 24, Type: typelabel.Ipaddr, Mandatory: true, Label: "TCPIP Netmask", Description: ""},
					smdx.PointElement{Id: DNS1_address, Offset: 26, Type: typelabel.Ipaddr, Mandatory: true, Label: "TCPIP DNS1", Description: ""},
					smdx.PointElement{Id: DNS2_address, Offset: 28, Type: typelabel.Ipaddr, Mandatory: true, Label: "TCPIP DNS2", Description: ""},
					smdx.PointElement{Id: Modbus_port, Offset: 30, Type: typelabel.Uint16, Mandatory: true, Label: "ModBus Port", Description: ""},
					smdx.PointElement{Id: SMTP_server_nm, Offset: 31, Type: typelabel.String, Length: 20, Mandatory: true, Label: "SMTP Server Name", Description: ""},
					smdx.PointElement{Id: SMTP_account_nm, Offset: 51, Type: typelabel.String, Length: 16, Mandatory: true, Label: "SMTP Account Name", Description: ""},
					smdx.PointElement{Id: SMTP_enable_SSL, Offset: 67, Type: typelabel.Enum16, Mandatory: true, Label: "Enable SMTP SSL", Description: ""},
					smdx.PointElement{Id: SMTP_password, Offset: 68, Type: typelabel.String, Length: 8, Mandatory: true, Label: "SMTP Password", Description: ""},
					smdx.PointElement{Id: SMTP_user_nm, Offset: 76, Type: typelabel.String, Length: 20, Mandatory: true, Label: "SMTP User Name", Description: ""},
					smdx.PointElement{Id: Stat_email_int, Offset: 96, Type: typelabel.Uint16, Mandatory: true, Label: "Status Email Interval", Description: ""},
					smdx.PointElement{Id: Stat_start_HR, Offset: 97, Type: typelabel.Uint16, Mandatory: true, Label: "Status Email Start Hour", Description: ""},
					smdx.PointElement{Id: Stat_email_sub, Offset: 98, Type: typelabel.String, Length: 25, Mandatory: true, Label: "Status Email Subject", Description: ""},
					smdx.PointElement{Id: Stat_email_addr1, Offset: 123, Type: typelabel.String, Length: 20, Mandatory: true, Label: "Status Email to Address 1", Description: ""},
					smdx.PointElement{Id: Stat_email_addr2, Offset: 143, Type: typelabel.String, Length: 20, Mandatory: true, Label: "Status Email to Address 2", Description: ""},
					smdx.PointElement{Id: Alarm_email_en, Offset: 163, Type: typelabel.Enum16, Mandatory: true, Label: "Enable Alarm Email", Description: ""},
					smdx.PointElement{Id: Alarm_email_sub, Offset: 164, Type: typelabel.String, Length: 25, Mandatory: true, Label: "Alarm Email Subject", Description: ""},
					smdx.PointElement{Id: Alarm_email_addr1, Offset: 189, Type: typelabel.String, Length: 20, Mandatory: true, Label: "Alarm Email to Address 1", Description: ""},
					smdx.PointElement{Id: Alarm_email_addr2, Offset: 209, Type: typelabel.String, Length: 20, Mandatory: true, Label: "Alarm Email to Address 2", Description: ""},
					smdx.PointElement{Id: FTP_password, Offset: 229, Type: typelabel.String, Length: 8, Mandatory: true, Label: "FTP Password", Description: ""},
					smdx.PointElement{Id: TELNET_password, Offset: 237, Type: typelabel.String, Length: 8, Mandatory: true, Label: "Telnet Password", Description: ""},
					smdx.PointElement{Id: Log_write_int, Offset: 245, Type: typelabel.Uint16, Units: "Tms", Mandatory: true, Label: "SD-Card Datalog Write Interval", Description: ""},
					smdx.PointElement{Id: Log_retain, Offset: 246, Type: typelabel.Uint16, Units: "Tmd", Mandatory: true, Label: "SD-Card Datalog Retain", Description: ""},
					smdx.PointElement{Id: Log_mode, Offset: 247, Type: typelabel.Enum16, Mandatory: true, Label: "SD-Card Datalog Mode", Description: ""},
					smdx.PointElement{Id: NTP_server_nm, Offset: 248, Type: typelabel.String, Length: 20, Mandatory: true, Label: "NTP Timer Server Name", Description: ""},
					smdx.PointElement{Id: NTP_enable, Offset: 268, Type: typelabel.Enum16, Mandatory: true, Label: "Enable Network Time", Description: ""},
					smdx.PointElement{Id: TimeZone, Offset: 269, Type: typelabel.Int16, Units: "Tmh", Mandatory: true, Label: "Time Zone", Description: ""},
					smdx.PointElement{Id: Date_year, Offset: 270, Type: typelabel.Uint16, Mandatory: true, Label: "Year", Description: ""},
					smdx.PointElement{Id: Date_month, Offset: 271, Type: typelabel.Uint16, Mandatory: true, Label: "Month", Description: ""},
					smdx.PointElement{Id: Date_Day, Offset: 272, Type: typelabel.Uint16, Mandatory: true, Label: "Day", Description: ""},
					smdx.PointElement{Id: Time_hour, Offset: 273, Type: typelabel.Uint16, Mandatory: true, Label: "Hour", Description: ""},
					smdx.PointElement{Id: Time_minute, Offset: 274, Type: typelabel.Uint16, Mandatory: true, Label: "Minute", Description: ""},
					smdx.PointElement{Id: Time_second, Offset: 275, Type: typelabel.Uint16, Mandatory: true, Label: "Second", Description: ""},
					smdx.PointElement{Id: Battery_temp, Offset: 276, Type: typelabel.Int16, ScaleFactor: "Temp_SF", Units: "C", Mandatory: true, Label: "Battery Temperature", Description: ""},
					smdx.PointElement{Id: Ambient_temp, Offset: 277, Type: typelabel.Int16, ScaleFactor: "Temp_SF", Units: "C", Mandatory: true, Label: "Ambient Temperature", Description: ""},
					smdx.PointElement{Id: Temp_SF, Offset: 278, Type: typelabel.ScaleFactor, Mandatory: true},
					smdx.PointElement{Id: AXS_Error, Offset: 279, Type: typelabel.Bitfield16, Mandatory: true, Label: "AXS Error", Description: ""},
					smdx.PointElement{Id: AXS_Status, Offset: 280, Type: typelabel.Bitfield16, Mandatory: true, Label: "AXS Status", Description: ""},
					smdx.PointElement{Id: AXS_Spare, Offset: 281, Type: typelabel.Uint16, Mandatory: true, Label: "Spare", Description: ""},
				},
			},
		}})
}
