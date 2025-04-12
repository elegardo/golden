package core

import (
	"reflect"
	"strings"

	"github.com/elegardo/golden/core/interfaces"
	"github.com/elegardo/golden/core/domain"
)

type RuleEvaluator struct {
	Comparator interfaces.Comparator
}

func (e *RuleEvaluator) Evaluate(operator domain.Operator, fact, value any) bool {
	switch operator {
	case domain.CO:
		return e.Comparator.Contains(fact, value)
	case domain.NC:
		return !e.Comparator.Contains(fact, value)
	case domain.IN:
		// TODO: print warning
		if theyAreNotString(fact, value) {
			return false
		}
		return strings.Contains(any(fact).(string), any(value).(string))
	case domain.NI:
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

func (e *RuleEvaluator) compare(operator domain.Operator, fact, value any) bool {
	result := e.Comparator.Compare(fact, value)
	switch operator {
	case domain.EQ:
		return result == 0
	case domain.NE:
		return result != 0
	case domain.GT:
		return result > 0
	case domain.GE:
		return result >= 0
	case domain.LT:
		return result < 0
	case domain.LE:
		return result <= 0
	default:
		panic("unsupported comparator")
	}
}

func theyAreNotString(fact, value any) bool {
	return reflect.TypeOf(fact).String() != "string" || reflect.TypeOf(value).String() != "string"
}
