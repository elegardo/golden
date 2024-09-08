package core

type Rule struct {
	Id          any
	Description string
	Event       Emmiter
	Conditions  []Condition
}

type Condition struct {
	Gate         Gate
	Conditionals []Conditional
}

type Conditional struct {
	Fact     string
	Operator Operator
	Value    any
}

// Logical Gates
// all	All conditions of rules are true
// any	Any conditions of rules are true
// none	None conditions of rules are true

type Gate int

const (
	ALL Gate = iota
	ANY
	NONE
)

func (g Gate) String() string {
	return [...]string{"all", "any", "none"}[g]
}

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

type Callback func(Emmiter)
