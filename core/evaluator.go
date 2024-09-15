package core

import (
	"reflect"
	"strings"

	"github.com/elegardo/golden/core/models"
)

type Evaluator struct {
}

func (e *Evaluator) Evaluate(operator models.Operator, fact, value any) bool {
	switch operator {
	case models.CO:
		return Contains(fact, value)
	case models.NC:
		return !Contains(fact, value)
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
		return e.compare(operator, fact, value)
	}
}

func (e *Evaluator) compare(operator models.Operator, fact, value any) bool {
	result := Compare(fact, value)
	switch operator {
	case models.EQ:
		return result == 0
	case models.NE:
		return result != 0
	case models.GT:
		return result > 0
	case models.GE:
		return result >= 0
	case models.LT:
		return result < 0
	case models.LE:
		return result <= 0
	default:
		panic("unsupported comparator")
	}
}

func theyAreNotString(fact, value any) bool {
	return reflect.TypeOf(fact).String() != "string" || reflect.TypeOf(value).String() != "string"
}
