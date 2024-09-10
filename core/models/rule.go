package models

type Rule struct {
	Id         any
	Priority   int
	Name       string
	Event      Emmiter
	Conditions []Condition
}
