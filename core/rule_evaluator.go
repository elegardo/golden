package core

type RuleEvaluator struct {
	Evaluator Evaluable
}

func (re *RuleEvaluator) Execute(ch chan<- bool, facts map[string]any, conditionals *[]Conditional) {
	for _, conditional := range *conditionals {
		if fact, exists := facts[conditional.Fact]; exists {
			ch <- re.evaluate(&conditional.Operator, fact, conditional.Value)
		}
	}
	close(ch)
}

func (re *RuleEvaluator) evaluate(operator *Operator, fact, value any) bool {
	return re.Evaluator.Evaluate(operator, fact, value)
}
