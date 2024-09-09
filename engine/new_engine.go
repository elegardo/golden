package engine

import "github.com/elegardo/golden/core"

// TODO: implement CLI flags
// -worker=10

func NewEngine(args ...string) *core.Engine {
	return &core.Engine{
		Match: &core.RuleMatch{
			RuleEvaluator: &core.RuleEvaluator{
				Evaluator: &core.Evaluator{},
			},
		},
	}
}
