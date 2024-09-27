package core

import (
	"github.com/elegardo/golden/core/interfaces"
	"github.com/elegardo/golden/core/models"
)

type Matcher struct {
	Evaluator interfaces.Evaluable
}

func (rm *Matcher) Execute(gate models.Gate, facts map[string]any, conditionals []models.Conditional) bool {
	switch gate {
	case models.ALL:
		return rm.allTrue(facts, conditionals)
	case models.ANY:
		return rm.anyTrue(facts, conditionals)
	case models.NONE:
		return rm.noneTrue(facts, conditionals)
	default:
		return false
	}
}

func (rm *Matcher) allTrue(facts map[string]any, conditionals []models.Conditional) bool {

	allTrue := false

	for _, conditional := range conditionals {
		if fact, exists := facts[conditional.Fact]; exists {
			if rm.Evaluator.Evaluate(conditional.Operator, fact, conditional.Value) {
				allTrue = true
			} else {
				allTrue = false
				break
			}
		}
	}

	return allTrue
}

func (rm *Matcher) anyTrue(facts map[string]any, conditionals []models.Conditional) bool {

	anyTrue := false

	for _, conditional := range conditionals {
		if fact, exists := facts[conditional.Fact]; exists {
			if rm.Evaluator.Evaluate(conditional.Operator, fact, conditional.Value) {
				anyTrue = true
				break
			}
		}
	}

	return anyTrue
}

func (rm *Matcher) noneTrue(facts map[string]any, conditionals []models.Conditional) bool {

	noneTrue := false

	for _, conditional := range conditionals {
		if fact, exists := facts[conditional.Fact]; exists {
			if rm.Evaluator.Evaluate(conditional.Operator, fact, conditional.Value) {
				noneTrue = false
				break
			} else {
				noneTrue = true
			}
		}
	}

	return noneTrue
}
