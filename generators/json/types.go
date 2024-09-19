package main

import (
	"encoding/json"
	"fmt"
	"io"
	"strconv"

	"github.com/andig/gosunspec/typelabel"
	"github.com/andig/gosunspec/typelen"
)

type Model struct {
	Group  Group
	Id     int
	Length int // TODO template only
}

type Group struct {
	Description string `json:"desc"`
	Label       string
	Name        string
	Points      []*Point
	Groups      []*SubGroup
	Id          int // TODO template only
	Length      int // TODO template only
}

type SubGroup struct {
	Description string `json:"desc"`
	Label       string
	Name        string
	Points      []*Point
	Length      int // TODO template only
}

type Point struct {
	Access      string
	Description string `json:"desc"`
	Label       string
	Mandatory   string
	Name        string
	Length      int         `json:"size"`
	ScaleFactor ScaleFactor `json:"sf"`
	Static      string
	Type        string
	Units       string
	Value       int
	Offset      int // TODO template only
}

type ScaleFactor string

func (s *ScaleFactor) UnmarshalJSON(b []byte) error {
	var sf string
	if err := json.Unmarshal(b, &sf); err != nil {
		var i int
		if err := json.Unmarshal(b, &i); err != nil {
			return err
		}
		sf = strconv.Itoa(i)
	}

	*s = ScaleFactor(sf)
	return nil
}

func FromJSON(reader io.Reader) (model Model, err error) {
	decoder := json.NewDecoder(reader)
	err = decoder.Decode(&model)
	return
}

func (m *Model) offsets(points []*Point) int {
	var offset int
	for _, p := range points {
		// if p.Name == "ID" || p.Name == "L" {
		// 	continue
		// }

		p.Offset = offset

		naturalLength := typelen.Length(p.Type)

		if p.Length == 0 {
			p.Length = int(naturalLength)

		} else if p.Type != typelabel.String && p.Length != int(naturalLength) {
			panic("length mismatch: " + fmt.Sprintf("%+v", p))
		}

		offset += p.Length
	}

	return offset
}

func filteredPoints(points []*Point) (res []*Point) {
	for _, p := range points {
		if p.Name != "ID" && p.Name != "L" {
			res = append(res, p)
		}
	}
	return
}

func (m *Model) Prepare() {
	m.Group.Id = m.Id
	m.Group.Length = m.offsets(filteredPoints(m.Group.Points))

	var length int
	for _, b := range m.Group.Groups {
		blockLen := m.offsets(b.Points)
		b.Length = blockLen
		length += blockLen
	}

	m.Length = m.Group.Length + length
}
