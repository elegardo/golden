package core

import (
	"github.com/elegardo/golden/core/interfaces"
	"github.com/elegardo/golden/core/models"
)

type RuleEvaluator struct {
	Evaluator interfaces.Evaluable
}

func (re *RuleEvaluator) Execute(ch chan<- bool, facts map[string]any, conditionals *[]models.Conditional) {
	for _, conditional := range *conditionals {
		if fact, exists := facts[conditional.Fact]; exists {
			ch <- re.evaluate(&conditional.Operator, fact, conditional.Value)
		}
	}
	//TODO: cerrar channel en un nivel superior (buena práctica)
	close(ch)
}

func (re *RuleEvaluator) evaluate(operator *models.Operator, fact, value any) bool {
	return re.Evaluator.Evaluate(operator, fact, value)
}
