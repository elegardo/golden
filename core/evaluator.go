package core

import "reflect"

type Evaluator struct {
}

func (e *Evaluator) Evaluate(operator *Operator, fact, value any) bool {
	switch *operator {
	case CO:
		a := Container[any]{Value: value}
		b := Container[any]{Value: fact}
		return a.Contains(b)
	case NC:
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

func (e *Evaluator) compare(operator *Operator, fact, value Comparable) bool {
	switch *operator {
	case EQ:
		return fact.Compare(value) == 0
	case NE:
		return fact.Compare(value) != 0
	case GT:
		return fact.Compare(value) > 0
	case GE:
		return fact.Compare(value) >= 0
	case LT:
		return fact.Compare(value) < 0
	case LE:
		return fact.Compare(value) <= 0
	default:
		panic("unsupported comparator")
	}
}
