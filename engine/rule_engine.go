package engine

// TODO: implement CLI flags
// -worker=10

func NewRuleEngine(args ...string) *Engine {
	return &Engine{
		Match: &RuleMatch{
			RuleEvaluator: &RuleEvaluator{
				Evaluator: &Evaluator{},
			},
		},
	}
}
