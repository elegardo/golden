package core

import (
	"github.com/elegardo/golden/core/interfaces"
	"github.com/elegardo/golden/core/models"
)

type Matcher struct {
	Evaluator interfaces.Evaluable
}

// TODO: Para ALL y NONE retornar lo mas pronto posible el false y terminar la ejecución y
// no esperar terminar de evaluar todas las conditions.
func (rm *Matcher) AllTrue(facts map[string]any, conditionals *[]models.Conditional) bool {
	ch := make(chan bool, len(*conditionals))

	go func() {
		for _, conditional := range *conditionals {
			if fact, exists := facts[conditional.Fact]; exists {
				if !rm.Evaluator.Evaluate(&conditional.Operator, fact, conditional.Value) {
					ch <- false
					break
				}
				ch <- true
			}
		}
		close(ch)
	}()

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

func (rm *Matcher) AnyTrue(facts map[string]any, conditionals *[]models.Conditional) bool {
	ch := make(chan bool, len(*conditionals))

	go func() {
		for _, conditional := range *conditionals {
			if fact, exists := facts[conditional.Fact]; exists {
				ch <- rm.Evaluator.Evaluate(&conditional.Operator, fact, conditional.Value)
			}
		}
		close(ch)
	}()

	anyTrue := false
	for result := range ch {
		if result {
			anyTrue = true
			break
		}
	}

	return anyTrue
}

func (rm *Matcher) NoneTrue(facts map[string]any, conditionals *[]models.Conditional) bool {
	ch := make(chan bool, len(*conditionals))

	for _, conditional := range *conditionals {
		if fact, exists := facts[conditional.Fact]; exists {
			ch <- rm.Evaluator.Evaluate(&conditional.Operator, fact, conditional.Value)
		}
	}

	close(ch)

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
