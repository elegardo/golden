package models

type Event interface {
	Order() int
	Value() any
}
