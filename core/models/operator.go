package models

// Logical Operators
// eq	Equal
// ne	Not equal
// gt	Greater than
// ge	Greater than or equal
// lt	Less than
// le	Less than or equal
// co	Contains (array)
// nc	No contains (array)
// in	String in other string
// ni	String not in other string

type Operator byte

const (
	EQ Operator = iota
	NE
	GT
	GE
	LT
	LE
	CO
	NC
	IN
	NI
)

func (o Operator) String() string {
	return [...]string{"eq", "ne", "gt", "ge", "lt", "le", "co", "nc", "in", "ni"}[o]
}
