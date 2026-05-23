// Package types holds the canonical, encoding-neutral representation
// of a SunSpec model definition. Concrete loaders for the SMDX/XML and
// JSON spec formats live in the types/xml and types/json sub-packages
// and produce values of these types directly.
//
// The types intentionally carry no encoding tags. Per-encoding details
// (XML attribute names, JSON polymorphic scalars, the SMDX "<strings>"
// indirection, the synthetic ID/L header registers) are normalised at
// the loader boundary.
package types

// BlockType classifies a block as either a single fixed-layout prefix
// or a repeating record. SunSpec models contain at most one fixed
// block followed by zero or more repeating blocks.
type BlockType string

const (
	// BlockFixed is the (optional) non-repeating prefix.
	BlockFixed BlockType = "fixed"
	// BlockRepeating is a repeating record.
	BlockRepeating BlockType = "repeating"
)

// Access is a point's read/write capability. Empty (the default)
// means read-only.
type Access string

const (
	AccessRead      Access = ""
	AccessReadWrite Access = "rw"
	AccessWriteOnly Access = "wo"
)

// Model is a canonical SunSpec model definition. It is identical
// whether sourced from SMDX XML or the JSON spec.
type Model struct {
	Id          uint16
	Name        string
	Label       string
	Description string
	Notes       string
	// Length is the on-wire size in registers (16-bit words),
	// excluding the synthetic ID and L header registers.
	Length uint16
	Blocks []Block
}

// Block is one fixed or repeating section of a model.
type Block struct {
	Name   string
	Type   BlockType
	Length uint16
	Points []Point
}

// Point is a single addressable field within a block.
type Point struct {
	Id          string
	Label       string
	Description string
	Notes       string
	Offset      uint16
	// Length is the register count for variable-width types
	// (e.g. strings); zero for fixed-width scalars.
	Length      uint16
	Type        string
	ScaleFactor string
	Units       string
	Mandatory   bool
	Access      Access
	Symbols     []Symbol
}

// Symbol is a named enum or bitfield member of a point.
type Symbol struct {
	Id    string
	Value string
}
