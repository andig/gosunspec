// Package xml loads SunSpec model definitions from the legacy SMDX
// XML format (spec/smdx/*.xml).
//
// All XML-specific tagging lives in this package; consumers receive a
// types.Model that has been normalised to the canonical shape.
package xml

import (
	"encoding/xml"
	"io"

	"github.com/andig/gosunspec/types"
)

// xmlModelDef is the root <sunSpecModels> element.
type xmlModelDef struct {
	XMLName xml.Name     `xml:"sunSpecModels"`
	Version string       `xml:"v,attr"`
	Models  []xmlModel   `xml:"model"`
	Strings []xmlStrings `xml:"strings"`
}

type xmlModel struct {
	Id     uint16     `xml:"id,attr"`
	Name   string     `xml:"name,attr"`
	Length uint16     `xml:"len,attr"`
	Blocks []xmlBlock `xml:"block"`
}

type xmlBlock struct {
	Name   string     `xml:"name,attr"`
	Length uint16     `xml:"len,attr"`
	Type   string     `xml:"type,attr"`
	Points []xmlPoint `xml:"point"`
}

type xmlPoint struct {
	Id          string      `xml:"id,attr"`
	Offset      uint16      `xml:"offset,attr"`
	Length      uint16      `xml:"len,attr"`
	Type        string      `xml:"type,attr"`
	ScaleFactor string      `xml:"sf,attr"`
	Units       string      `xml:"units,attr"`
	Mandatory   bool        `xml:"mandatory,attr"`
	Access      string      `xml:"access,attr"`
	Symbols     []xmlSymbol `xml:"symbol"`
}

type xmlSymbol struct {
	Id    string `xml:"id,attr"`
	Value string `xml:",chardata"`
}

type xmlStrings struct {
	Id           string             `xml:"id,attr"`
	Locale       string             `xml:"locale,attr"`
	ModelStrings xmlModelStrings    `xml:"model"`
	PointStrings []xmlPointStrings  `xml:"point"`
}

type xmlModelStrings struct {
	Label       string `xml:"label"`
	Description string `xml:"description"`
	Notes       string `xml:"notes"`
}

type xmlPointStrings struct {
	Id          string `xml:"id,attr"`
	Label       string `xml:"label"`
	Description string `xml:"description"`
	Notes       string `xml:"notes"`
}

// Decode parses one SMDX XML file and returns the first model contained
// in it. SunSpec model files conventionally contain a single model; if
// the file contains more, only the first is returned (the others are
// silently ignored to match historical behaviour).
//
// Strings information (model label/description, point labels) is taken
// from the first <strings> element, regardless of locale.
func Decode(r io.Reader) (*types.Model, error) {
	var doc xmlModelDef
	if err := xml.NewDecoder(r).Decode(&doc); err != nil {
		return nil, err
	}
	if len(doc.Models) == 0 {
		return nil, nil
	}
	return convert(&doc), nil
}

func convert(doc *xmlModelDef) *types.Model {
	xm := doc.Models[0]

	var ms xmlModelStrings
	pointLabels := map[string]xmlPointStrings{}
	if len(doc.Strings) > 0 {
		ms = doc.Strings[0].ModelStrings
		for _, ps := range doc.Strings[0].PointStrings {
			pointLabels[ps.Id] = ps
		}
	}

	m := &types.Model{
		Id:          xm.Id,
		Name:        xm.Name,
		Length:      xm.Length,
		Label:       ms.Label,
		Description: ms.Description,
		Notes:       ms.Notes,
		Blocks:      make([]types.Block, 0, len(xm.Blocks)),
	}

	for _, xb := range xm.Blocks {
		bt := types.BlockType(xb.Type)
		if bt == "" {
			bt = types.BlockFixed
		}
		b := types.Block{
			Name:   xb.Name,
			Type:   bt,
			Length: xb.Length,
			Points: make([]types.Point, 0, len(xb.Points)),
		}
		for _, xp := range xb.Points {
			ps := pointLabels[xp.Id]
			b.Points = append(b.Points, types.Point{
				Id:          xp.Id,
				Label:       ps.Label,
				Description: ps.Description,
				Notes:       ps.Notes,
				Offset:      xp.Offset,
				Length:      xp.Length,
				Type:        xp.Type,
				ScaleFactor: xp.ScaleFactor,
				Units:       xp.Units,
				Mandatory:   xp.Mandatory,
				Access:      types.Access(xp.Access),
				Symbols:     convertSymbols(xp.Symbols),
			})
		}
		m.Blocks = append(m.Blocks, b)
	}
	return m
}

func convertSymbols(ss []xmlSymbol) []types.Symbol {
	if len(ss) == 0 {
		return nil
	}
	out := make([]types.Symbol, 0, len(ss))
	for _, s := range ss {
		out = append(out, types.Symbol{Id: s.Id, Value: s.Value})
	}
	return out
}
