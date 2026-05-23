// Package json loads SunSpec model definitions from the JSON spec
// format (spec/json/*.json published at github.com/sunspec/models).
//
// All JSON-specific tagging lives in this package; consumers receive a
// types.Model that has been normalised to the canonical shape.
package json

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/andig/gosunspec/types"
)

type jsonModel struct {
	Id    uint16    `json:"id"`
	Group jsonGroup `json:"group"`
}

type jsonGroup struct {
	Name   string      `json:"name"`
	Label  string      `json:"label"`
	Desc   string      `json:"desc"`
	Detail string      `json:"detail"`
	Notes  string      `json:"notes"`
	Type   string      `json:"type"`
	Count  interface{} `json:"count,omitempty"`
	Points []jsonPoint `json:"points"`
	Groups []jsonGroup `json:"groups"`
}

type jsonPoint struct {
	Name      string       `json:"name"`
	Label     string       `json:"label"`
	Desc      string       `json:"desc"`
	Detail    string       `json:"detail"`
	Notes     string       `json:"notes"`
	Type      string       `json:"type"`
	Size      uint16       `json:"size"`
	Mandatory string       `json:"mandatory"`
	Access    string       `json:"access"`
	// SF is usually a reference to another point by name ("A_SF"), but
	// some models declare a literal integer power-of-ten (e.g. -1).
	SF        interface{}  `json:"sf"`
	Units     string       `json:"units"`
	Static    string       `json:"static"`
	Symbols   []jsonSymbol `json:"symbols"`
	Standards []string     `json:"standards"`
	Value     interface{}  `json:"value,omitempty"`
}

type jsonSymbol struct {
	Name  string      `json:"name"`
	Value interface{} `json:"value"`
	Label string      `json:"label,omitempty"`
	Desc  string      `json:"desc,omitempty"`
	Notes string      `json:"notes,omitempty"`
}

// Decode parses one SunSpec JSON model file and returns the normalised
// model definition.
//
// Conversion notes:
//   - The outer group's "ID" and "L" points are stripped (the synthetic
//     header registers are not part of any block).
//   - Point offsets are computed from the cumulative declared sizes.
//   - Nested groups become repeating blocks. Repeated names within a
//     model are disambiguated with a numeric suffix.
//   - Access is normalised to the canonical lowercase form.
//   - mandatory="M" becomes Mandatory: true.
func Decode(r io.Reader) (*types.Model, error) {
	var jm jsonModel
	dec := json.NewDecoder(r)
	dec.UseNumber()
	if err := dec.Decode(&jm); err != nil {
		return nil, fmt.Errorf("decode json model: %w", err)
	}
	return convert(jm), nil
}

func convert(jm jsonModel) *types.Model {
	fixedPoints, fixedLen := convertPoints(stripHeader(jm.Group.Points))

	m := &types.Model{
		Id:          jm.Id,
		Name:        jm.Group.Name,
		Label:       jm.Group.Label,
		Description: jm.Group.Desc,
		Notes:       jm.Group.Notes,
	}

	// Models like 304 (inclinometer) have only a repeating block; the
	// outer group is just the synthetic header. Omit the fixed block in
	// that case so the runtime length math matches the XML/SMDX shape.
	if len(fixedPoints) > 0 {
		m.Blocks = append(m.Blocks, types.Block{
			Type:   types.BlockFixed,
			Length: fixedLen,
			Points: fixedPoints,
		})
	}

	totalLen := fixedLen
	seen := map[string]int{}
	var walk func(groups []jsonGroup)
	walk = func(groups []jsonGroup) {
		for _, g := range groups {
			name := g.Name
			if seen[g.Name] > 0 {
				name = fmt.Sprintf("%s_%d", g.Name, seen[g.Name]+1)
			}
			seen[g.Name]++

			pts, blen := convertPoints(g.Points)
			m.Blocks = append(m.Blocks, types.Block{
				Name:   name,
				Type:   types.BlockRepeating,
				Length: blen,
				Points: pts,
			})
			totalLen += blen
			walk(g.Groups)
		}
	}
	walk(jm.Group.Groups)

	m.Length = totalLen
	return m
}

// stripHeader removes the synthetic ID and L points that JSON definitions
// include at the start of every model. They are header registers, not
// real block fields.
func stripHeader(points []jsonPoint) []jsonPoint {
	out := make([]jsonPoint, 0, len(points))
	for _, p := range points {
		if p.Name == "ID" || p.Name == "L" {
			continue
		}
		out = append(out, p)
	}
	return out
}

func convertPoints(points []jsonPoint) ([]types.Point, uint16) {
	out := make([]types.Point, 0, len(points))
	var offset uint16
	for _, p := range points {
		out = append(out, types.Point{
			Id:          p.Name,
			Label:       p.Label,
			Description: p.Desc,
			Notes:       p.Notes,
			Offset:      offset,
			Length:      pointLength(p),
			Type:        p.Type,
			ScaleFactor: stringify(p.SF),
			Units:       p.Units,
			Mandatory:   p.Mandatory == "M",
			Access:      normalizeAccess(p.Access),
			Symbols:     convertSymbols(p.Symbols),
		})
		offset += p.Size
	}
	return out, offset
}

// pointLength preserves the SMDX convention of emitting an explicit
// length only for variable-width types. For scalars the length is
// implicit in the type.
func pointLength(p jsonPoint) uint16 {
	if p.Type == "string" {
		return p.Size
	}
	return 0
}

func normalizeAccess(a string) types.Access {
	switch a {
	case "RW", "rw":
		return types.AccessReadWrite
	case "R", "r":
		return types.AccessRead
	case "WO", "wo":
		return types.AccessWriteOnly
	}
	return types.AccessRead
}

// stringify renders a polymorphic JSON scalar (string / json.Number /
// int / float64) as its canonical string form. Returns "" for nil so
// callers can use len(...) == 0 as the "absent" check.
func stringify(v interface{}) string {
	switch x := v.(type) {
	case nil:
		return ""
	case string:
		return x
	case json.Number:
		return x.String()
	case float64:
		i := int64(x)
		if float64(i) == x {
			return fmt.Sprintf("%d", i)
		}
		return fmt.Sprintf("%g", x)
	case bool:
		if x {
			return "true"
		}
		return "false"
	default:
		return fmt.Sprintf("%v", x)
	}
}

func convertSymbols(ss []jsonSymbol) []types.Symbol {
	if len(ss) == 0 {
		return nil
	}
	out := make([]types.Symbol, 0, len(ss))
	for _, s := range ss {
		out = append(out, types.Symbol{
			Id:    s.Name,
			Value: stringify(s.Value),
		})
	}
	return out
}
