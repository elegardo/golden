package core

import (
	"github.com/elegardo/golden/core/interfaces"
	"github.com/elegardo/golden/core/models"
)

type Matcher struct {
	Evaluator interfaces.Evaluable
}

// 4 allocs/op
func (rm *Matcher) AllTrue(facts map[string]any, conditionals []models.Conditional) bool {
	ch := make(chan bool, len(conditionals))

	go rm.evaluate(ch, facts, conditionals)

	allTrue := false
	for result := range ch {
		if result {
			allTrue = true
		} else {
			allTrue = false
			break
		}
	}

	return allTrue
}

// 4 allocs/op
func (rm *Matcher) AnyTrue(facts map[string]any, conditionals []models.Conditional) bool {
	ch := make(chan bool, len(conditionals))

	go rm.evaluate(ch, facts, conditionals)

	anyTrue := false
	for result := range ch {
		if result {
			anyTrue = true
			break
		}
	}

	return anyTrue
}

// 4 allocs/op
func (rm *Matcher) NoneTrue(facts map[string]any, conditionals []models.Conditional) bool {
	ch := make(chan bool, len(conditionals))

	go rm.evaluate(ch, facts, conditionals)

	noneTrue := false
	for result := range ch {
		if result {
			noneTrue = false
			break
		} else {
			noneTrue = true
		}
	}

	return noneTrue
}

func (rm *Matcher) evaluate(ch chan bool, facts map[string]any, conditionals []models.Conditional) {
	for _, conditional := range conditionals {
		if fact, exists := facts[conditional.Fact]; exists {
			ch <- rm.Evaluator.Evaluate(conditional.Operator, fact, conditional.Value)
		}
	}
	close(ch)
}
