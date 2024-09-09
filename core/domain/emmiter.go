package domain

type Emmiter interface {
	Order() int
	Value() any
}
