package impl

import (
	"log"

	sunspec "github.com/andig/gosunspec"
	"github.com/andig/gosunspec/types"
	"github.com/andig/gosunspec/spi"
	"github.com/andig/gosunspec/typelabel"
)

func NewArray() spi.ArraySPI {
	return &array{
		devices: []*device{},
	}
}

func NewDevice() spi.DeviceSPI {
	return &device{
		models: []*model{},
	}
}

// Answer a new model that can be mapped to a contiguous range of bytes. The size
// (or bound) of that range determines the number of repeats.
func NewContiguousModel(md *types.Model, bound uint16, phys spi.Driver) spi.ModelSPI {
	block0len := int(md.Blocks[0].Length) // specified length of the first block
	excess := int(bound) - block0len      // number of words for additional blocks
	nr := 0                               // number of repeat blocks
	rlen := block0len                     // repeat block length

	// reset the repeat block length if there is a repeat block

	if len(md.Blocks) > 1 {
		rlen = int(md.Blocks[1].Length)
	}

	// calculate the repeats and normalize negative numbers to zero

	nr = (excess / rlen)
	if nr < 0 {
		nr = 0
	}

	// build a model with the specified geometry

	m := NewModel(md, nr, phys)

	// handle a specification special case &/or report length inconsistenes.

	if len(md.Blocks) == 1 && excess < 0 {
		// Models with fixed-blocks only are allowed to truncate the
		// fixed block to a smaller model length because of a special
		// case for the Common Model allowed for by the SunSpec Information Model.
		//
		// (Note: it isn't clear why the specification needs to cater for
		// this special case.)
		//
		// In particular, implementations are allowed to return a Common
		// Model block of length 65 instead of 66 if it suits their purposes to do so.
		//
		// We support truncation of other fixed-blocks only models in the same
		// way, if it occurs (e.g. model length < block length) just in case i
		// the standard has been interpreted by others to permit this for
		// other models.
		//
		// The same support cannot be extended to models with repeating blocks
		// because the specification doesn't provide guidance about where to
		// apply such truncations in this case.
		m.MustBlock(0).(spi.BlockSPI).SetLength(bound)
	} else {
		// If the model length is not strictly identical to the canaonical
		// length of a model instance with 'nr' repeat blocks, then report
		// the inconsistency but don't otherwise do anything. It is likely
		// that API may not correctly map points of at least some blocks
		// in this case.
		if nr*rlen+block0len != int(bound) {
			log.Printf("warning: inconsistent model and block lengths: %d, %d, %d*%d. YMMV.", bound, block0len, nr, rlen)
		}
	}
	return m
}

// Answer a new model with nr repeats of the repating block, if any. The caller
// determines the number of repeats.
func NewModel(md *types.Model, nr int, driver spi.Driver) spi.ModelSPI {
	blocks := []*block{newBlock(&md.Blocks[0], driver, nil)}

	m := &model{
		def:    md,
		driver: driver,
		blocks: blocks,
	}

	for nr > 0 {
		nr--
		m.AddRepeat()
	}

	return m
}

// Answer a new block.
// The fixedBlock param is used to find scale factors that are
// outside of a repeating block.
func newBlock(blockDef *types.Block, driver spi.Driver, fixedBlock *block) *block {
	b := &block{
		def:    blockDef,
		points: map[string]*point{},
		driver: driver,
	}

	p := []*types.Point{}
	q := []*types.Point{}

	for x, pe := range blockDef.Points {
		if pe.Type == typelabel.ScaleFactor {
			p = append(p, &blockDef.Points[x])
		} else {
			q = append(q, &blockDef.Points[x])
		}
	}

	p = append(p, q...)

	for _, pe := range p {
		var sfp sunspec.Point
		var sp *point
		var ok bool

		if sfp, ok = b.points[pe.ScaleFactor]; ok {
			sp = newPoint(pe, sfp, b)
		} else if pe.ScaleFactor != "" && fixedBlock != nil {
			if sfp, ok = fixedBlock.points[pe.ScaleFactor]; ok {
				sp = newPoint(pe, sfp, b)
			} else {
				sp = newPoint(pe, nil, b)
			}
		} else {
			sp = newPoint(pe, nil, b)
		}

		b.points[pe.Id] = sp
	}

	return b
}

// Answer a new point.
func newPoint(pd *types.Point, scaleFactor sunspec.Point, b *block) *point {
	return &point{
		def:         pd,
		scaleFactor: scaleFactor,
		block:       b,
		err:         errNotInitialized,
	}
}
