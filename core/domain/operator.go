package domain

// Logical Operators
// eq	Equal
// ne	Not equal
// gt	Greater than
// ge	Greater than or equal
// lt	Less than
// le	Less than or equal
// co	Contains (string, array)
// nc	No contains (string, array)

type Operator int

const (
	EQ Operator = iota
	NE
	GT
	GE
	LT
	LE
	CO
	NC
)

func (o Operator) String() string {
	return [...]string{"eq", "ne", "gt", "ge", "lt", "le", "co", "nc"}[o]
}
