package core

import (
	"reflect"

	"github.com/elegardo/golden/core/domain"
	"github.com/elegardo/golden/core/interfaces"
)

type Evaluator struct {
}

func (e *Evaluator) Evaluate(operator *domain.Operator, fact, value any) bool {
	switch *operator {
	case domain.CO:
		a := Container[any]{Value: value}
		b := Container[any]{Value: fact}
		return a.Contains(b)
	case domain.NC:
		a := Container[any]{Value: value}
		b := Container[any]{Value: fact}
		return !a.Contains(b)
	default:
		// different types cannot be compared
		if reflect.TypeOf(fact) != reflect.TypeOf(value) {
			return false
		}
		return e.compare(operator, Comparator[any]{Value: fact}, Comparator[any]{Value: value})
	}
}

func (e *Evaluator) compare(operator *domain.Operator, fact, value interfaces.Comparable) bool {
	switch *operator {
	case domain.EQ:
		return fact.Compare(value) == 0
	case domain.NE:
		return fact.Compare(value) != 0
	case domain.GT:
		return fact.Compare(value) > 0
	case domain.GE:
		return fact.Compare(value) >= 0
	case domain.LT:
		return fact.Compare(value) < 0
	case domain.LE:
		return fact.Compare(value) <= 0
	default:
		panic("unsupported comparator")
	}
}
