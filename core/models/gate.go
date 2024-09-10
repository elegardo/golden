package models

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
