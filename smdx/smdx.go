package smdx

import (
	"encoding/xml"
	"io"
)

type ModelDefinitionElement struct {
	XMLName xml.Name         `xml:"sunSpecModels"`
	Version string           `xml:"v,attr"`
	Models  []ModelElement   `xml:"model"`
	Strings []StringsElement `xml:"strings"`
}

type ModelElement struct {
	XMLName xml.Name       `xml:"model"`
	Id      uint16         `xml:"id,attr"`
	Name    string         `xml:"name,attr"`
	Length  uint16         `xml:"len,attr"`
	Blocks  []BlockElement `xml:"block"`
}

type BlockElement struct {
	XMLName xml.Name       `xml:"block"`
	Name    string         `xml:"name,attr"`
	Length  uint16         `xml:"len,attr"`
	Type    string         `xml:"type,attr"`
	Points  []PointElement `xml:"point"`
}

type PointElement struct {
	XMLName     xml.Name        `xml:"point"`
	Id          string          `xml:"id,attr"`
	Label       string          `xml:",omit"`
	Description string          `xml:",omit"`
	Offset      uint16          `xml:"offset,attr"`
	Length      uint16          `xml:"len,attr"`
	Type        string          `xml:"type,attr"`
	ScaleFactor string          `xml:"sf,attr"`
	Units       string          `xml:"units,attr"`
	Mandatory   bool            `xml:"mandatory,attr"`
	Access      string          `xml:"access,attr"`
	Symbols     []SymbolElement `xml:"symbol"`
}

type SymbolElement struct {
	XMLName xml.Name `xml:"symbol"`
	Id      string   `xml:"id,attr"`
	Value   string   `xml:",chardata"`
}

type StringsElement struct {
	XMLName      xml.Name              `xml:"strings"`
	Id           string                `xml:"id,attr"`
	Locale       string                `xml:"locale,attr"`
	ModelStrings ModelStringsElement   `xml:"model"`
	PointStrings []PointStringsElement `xml:"point"`
}

type ModelStringsElement struct {
	XMLName     xml.Name `xml:"model"`
	Label       string   `xml:"label"`
	Description string   `xml:"description"`
	Notes       string   `xml:"notes"`
}

type PointStringsElement struct {
	XMLName     xml.Name `xml:"point"`
	Id          string   `xml:"id,attr"`
	Label       string   `xml:"label"`
	Description string   `xml:"description"`
	Notes       string   `xml:"notes"`
}

func FromXML(reader io.Reader) (model ModelDefinitionElement, err error) {
	decoder := xml.NewDecoder(reader)
	err = decoder.Decode(&model)
	return
}
