package models

type Rule struct {
	Id          any
	Description string
	Event       Emmiter
	Conditions  []Condition
}
