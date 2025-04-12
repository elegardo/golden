package domain

type Rule struct {
	Id         any
	Priority   int
	Name       string
	Event      Event
	Conditions []Condition
}
