package domain

type Event interface {
	Order() int
	Value() any
}
