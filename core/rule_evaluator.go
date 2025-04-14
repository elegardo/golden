package core

import (
	"strings"

	"github.com/elegardo/golden/core/interfaces"
	"github.com/elegardo/golden/core/domain"
)

type RuleEvaluator struct {
	Comparator interfaces.Comparator
}

func (e *RuleEvaluator) Evaluate(operator domain.Operator, pair *domain.Pair) bool {
	switch operator {
	case domain.CO:
		return e.Comparator.Contains(pair)
	case domain.NC:
		return !e.Comparator.Contains(pair)
	case domain.IN:
		// TODO: print warning
		if !pair.AreString() {
			return false
		}
		return strings.Contains(pair.GetString1(), pair.GetString2())
	case domain.NI:
		// TODO: print warning
		if !pair.AreString() {
			return false
		}
		return !strings.Contains(pair.GetString1(), pair.GetString2())
	default:
		// different types cannot be compared
		// TODO: print warning
		if !pair.AreSameType() {
			return false
		}
		return e.compare(operator, pair)
	}
}

func (e *RuleEvaluator) compare(operator domain.Operator, pair *domain.Pair) bool {
	result := e.Comparator.Compare(pair)
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
