package engine

import (
	. "github.com/elegardo/golden/core"
)

// TODO: implement CLI flags
// -worker=10

func NewEngine(args ...string) *Engine {
	return &Engine{
		Match: &RuleMatch{
			RuleEvaluator: &RuleEvaluator{
				Evaluator: &Evaluator{},
			},
		},
	}
}
