package models

type Conditional struct {
	Fact     string
	Operator Operator
	Value    any
}
