package core

import "github.com/elegardo/golden/core/models"

type RuleMatch struct {
	RuleEvaluator *RuleEvaluator
}

func (rm *RuleMatch) AllTrue(main chan<- bool, facts map[string]any, conditionals *[]models.Conditional) {
	ch := make(chan bool)
	go rm.RuleEvaluator.Execute(ch, facts, conditionals)

	allTrue := false
	for result := range ch {
		if result {
			allTrue = true
		} else {
			allTrue = false
			break
		}
	}

	main <- allTrue
}

func (rm *RuleMatch) AnyTrue(main chan<- bool, facts map[string]any, conditionals *[]models.Conditional) {
	ch := make(chan bool)
	go rm.RuleEvaluator.Execute(ch, facts, conditionals)

	anyTrue := false
	for result := range ch {
		if result {
			anyTrue = true
			break
		}
	}

	main <- anyTrue
}

func (rm *RuleMatch) NoneTrue(main chan<- bool, facts map[string]any, conditionals *[]models.Conditional) {
	ch := make(chan bool)
	go rm.RuleEvaluator.Execute(ch, facts, conditionals)

	noneTrue := false
	for result := range ch {
		if result {
			noneTrue = false
			break
		} else {
			noneTrue = true
		}
	}

	main <- noneTrue
}
