package impl

import (
	sunspec "github.com/andig/gosunspec"
	"github.com/andig/gosunspec/smdx"
	"github.com/andig/gosunspec/spi"
)

type model struct {
	anchored
	driver spi.Driver
	smdx   *smdx.ModelElement
	blocks []*block
}

func (m *model) Id() sunspec.ModelId {
	return sunspec.ModelId(m.smdx.Id)
}

func (m *model) Blocks() int {
	return len(m.blocks)
}

func (m *model) Block(i int) (sunspec.Block, error) {
	if i < len(m.blocks) {
		return m.blocks[i], nil
	} else {
		return nil, sunspec.ErrNoSuchBlock
	}
}

func (m *model) MustBlock(i int) sunspec.Block {
	if b, err := m.Block(i); err != nil {
		panic(err)
	} else {
		return b
	}
}

func (m *model) Do(f func(b sunspec.Block)) {
	for _, b := range m.blocks {
		f(b)
	}
}

func (m *model) AddRepeat() error {
	fixed := m.blocks[0]
	repeat := &m.smdx.Blocks[len(m.smdx.Blocks)-1]
	m.blocks = append(m.blocks, newBlock(repeat, m.driver, fixed))
	return nil
}

func (m *model) Length() uint16 {
	l := uint16(0)
	for _, bl := range m.blocks {
		l += bl.Length()
	}
	return l
}
