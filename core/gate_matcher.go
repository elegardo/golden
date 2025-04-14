package core

import (
	"github.com/elegardo/golden/core/interfaces"
	"github.com/elegardo/golden/core/domain"
)

type GateMatcher struct {
	Evaluator interfaces.Evaluator
}

func (rm *GateMatcher) Match(gate domain.Gate, facts map[string]any, conditionals []domain.Conditional) bool {
	switch gate {
	case domain.ALL:
		return rm.allTrue(facts, conditionals)
	case domain.ANY:
		return rm.anyTrue(facts, conditionals)
	case domain.NONE:
		return rm.noneTrue(facts, conditionals)
	default:
		return false
	}
}

func (rm *GateMatcher) allTrue(facts map[string]any, conditionals []domain.Conditional) bool {

	allTrue := false

	for _, conditional := range conditionals {
		if factValue, exists := facts[conditional.Fact]; exists {
			pair := domain.NewPair(factValue, conditional.Value)
			if rm.Evaluator.Evaluate(conditional.Operator, pair) {
				allTrue = true
			} else {
				allTrue = false
				break
			}
		}
	}

	return allTrue
}

func (rm *GateMatcher) anyTrue(facts map[string]any, conditionals []domain.Conditional) bool {

	anyTrue := false

	for _, conditional := range conditionals {
		if factValue, exists := facts[conditional.Fact]; exists {
			pair := domain.NewPair(factValue, conditional.Value)
			if rm.Evaluator.Evaluate(conditional.Operator, pair) {
				anyTrue = true
				break
			}
		}
	}

	return anyTrue
}

func (rm *GateMatcher) noneTrue(facts map[string]any, conditionals []domain.Conditional) bool {

	noneTrue := false

	for _, conditional := range conditionals {
		if factValue, exists := facts[conditional.Fact]; exists {
			pair := domain.NewPair(factValue, conditional.Value)
			if rm.Evaluator.Evaluate(conditional.Operator, pair) {
				noneTrue = false
				break
			} else {
				noneTrue = true
			}
		}
	}

	return noneTrue
}
