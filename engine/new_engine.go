package engine

import (
	"github.com/elegardo/golden/core"
	"github.com/elegardo/golden/core/interfaces"
)

// TODO: implement CLI flags
// -workers-pool=10

func worker() interfaces.Workereable {
	return &core.Worker{
		Match: &core.RuleMatch{
			RuleEvaluator: &core.RuleEvaluator{
				Evaluator: &core.Evaluator{},
			},
		},
	}
}

func NewRunnerEngine(args ...string) interfaces.Runnable {
	return &RunnerEngine{
		Worker: worker(),
	}
}

func NewAsyncEngine(args ...string) interfaces.Engine {
	return &AsyncEngine{
		Worker: worker(),
	}
}
