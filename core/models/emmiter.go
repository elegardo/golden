package models

type Emmiter interface {
	Order() int
	Value() any
}
