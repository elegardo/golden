package core

import (
	"reflect"
	"strings"

	"github.com/elegardo/golden/core/interfaces"
	"github.com/elegardo/golden/core/models"
)

type Evaluator struct {
}

func (e *Evaluator) Evaluate(operator *models.Operator, fact, value any) bool {
	switch *operator {
	case models.CO:
		a := Container[any]{Value: value}
		b := Container[any]{Value: fact}
		return a.Contains(b)
	case models.NC:
		a := Container[any]{Value: value}
		b := Container[any]{Value: fact}
		return !a.Contains(b)
	case models.IN:
		// TODO: print warning
		if theyAreNotString(fact, value) {
			return false
		}
		return strings.Contains(any(fact).(string), any(value).(string))
	case models.NI:
		// TODO: print warning
		if theyAreNotString(fact, value) {
			return false
		}
		return !strings.Contains(any(fact).(string), any(value).(string))
	default:
		// different types cannot be compared
		// TODO: print warning
		if reflect.TypeOf(fact) != reflect.TypeOf(value) {
			return false
		}
		return e.compare(operator, Comparator[any]{Value: fact}, Comparator[any]{Value: value})
	}
}

func (e *Evaluator) compare(operator *models.Operator, fact, value interfaces.Comparable) bool {
	switch *operator {
	case models.EQ:
		return fact.Compare(value) == 0
	case models.NE:
		return fact.Compare(value) != 0
	case models.GT:
		return fact.Compare(value) > 0
	case models.GE:
		return fact.Compare(value) >= 0
	case models.LT:
		return fact.Compare(value) < 0
	case models.LE:
		return fact.Compare(value) <= 0
	default:
		panic("unsupported comparator")
	}
}

func theyAreNotString(fact, value any) bool {
	return reflect.TypeOf(fact).String() != "string" || reflect.TypeOf(value).String() != "string"
}
