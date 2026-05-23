package main

import (
	"fmt"
	"log"
	"os"

	_ "github.com/andig/gosunspec/models"
	"github.com/andig/gosunspec/typelabel"
	"github.com/andig/gosunspec/typelen"
	"github.com/andig/gosunspec/types"
)

type Type struct {
	Name   string
	Length int
}

var typeList = []Type{
	{typelabel.Acc16, typelen.Acc16},
	{typelabel.Acc32, typelen.Acc32},
	{typelabel.Acc64, typelen.Acc64},
	{typelabel.Bitfield16, typelen.Bitfield16},
	{typelabel.Bitfield32, typelen.Bitfield32},
	{typelabel.Count, typelen.Count},
	{typelabel.Enum16, typelen.Enum16},
	{typelabel.Enum32, typelen.Enum32},
	{typelabel.Eui48, typelen.Eui48},
	{typelabel.Float32, typelen.Float32},
	{typelabel.Int16, typelen.Int16},
	{typelabel.Int32, typelen.Int32},
	{typelabel.Int64, typelen.Int64},
	{typelabel.Ipaddr, typelen.Ipaddr},
	{typelabel.Ipv6addr, typelen.Ipv6addr},
	{typelabel.Pad, typelen.Pad},
	{typelabel.String, typelen.String},
	{typelabel.ScaleFactor, typelen.ScaleFactor},
	{typelabel.Uint16, typelen.Uint16},
	{typelabel.Uint32, typelen.Uint32},
	{typelabel.Uint64, typelen.Uint64},
}

var typeMap = map[string]Type{}

func init() {
	for _, v := range typeList {
		typeMap[v.Name] = v
	}
}

func main() {
	if err := types.DoModels(func(m *types.Model) error {
		if len(m.Blocks) < 1 {
			fmt.Fprintf(os.Stderr, "%d: not enough blocks\n", m.Id)
		} else if len(m.Blocks) > 2 {
			fmt.Fprintf(os.Stderr, "%d: too many blocks\n", m.Id)
		}
		totalBlockLen := 0
		for _, b := range m.Blocks {
			totalBlockLen += int(b.Length)
		}
		if totalBlockLen != int(m.Length) {
			fmt.Fprintf(os.Stderr, "%d: total block length mismatch\n", m.Id)
		}
		for bx, b := range m.Blocks {
			totalPointLen := 0
			for _, p := range b.Points {
				t := typeMap[p.Type]
				if t.Length == 0 {
					if p.Length == 0 {
						fmt.Fprintf(os.Stderr, "%d: %d: %s: zero length point\n", m.Id, bx, p.Id)
					}
					totalPointLen += int(p.Length)
				} else {
					totalPointLen += int(t.Length)
					if p.Length != 0 && int(p.Length) != t.Length {
						fmt.Fprintf(os.Stderr, "%d: %d: %s: inconsistent point length\n", m.Id, bx, p.Id)
					}
				}
			}
			if int(b.Length)-totalPointLen > 1 || int(b.Length) < totalPointLen {
				fmt.Fprintf(os.Stderr, "%d: %d: block length inconsistent with point length: %d, %d\n", m.Id, bx, b.Length, totalPointLen)

			}
		}
		return nil
	}); err != nil {
		log.Fatalf("failed: %v", err)
	}
}
