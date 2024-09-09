package core

import (
	"github.com/elegardo/golden/core/domain"
	"github.com/elegardo/golden/core/interfaces"
)

type RuleEvaluator struct {
	Evaluator interfaces.Evaluable
}

func (re *RuleEvaluator) Execute(ch chan<- bool, facts map[string]any, conditionals *[]domain.Conditional) {
	for _, conditional := range *conditionals {
		if fact, exists := facts[conditional.Fact]; exists {
			ch <- re.evaluate(&conditional.Operator, fact, conditional.Value)
		}
	}
	close(ch)
}

func (re *RuleEvaluator) evaluate(operator *domain.Operator, fact, value any) bool {
	return re.Evaluator.Evaluate(operator, fact, value)
}
