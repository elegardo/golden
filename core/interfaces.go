package core

type Comparable interface {
	Compare(Comparable) int
}

type Containable interface {
	Contains(Containable) bool
}

type Evaluable interface {
	Evaluate(operator *Operator, fact, value any) bool
}

type Emmiter interface {
	Order() int
	Value() any
}

type Matchable interface {
	allTrue(main chan<- bool, facts map[string]any, conditionals *[]Conditional)
	anyTrue(main chan<- bool, facts map[string]any, conditionals *[]Conditional)
	noneTrue(main chan<- bool, facts map[string]any, conditionals *[]Conditional)
}
