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

func NewSyncEngine(args ...string) interfaces.Engineable {
	return &SyncEngine{
		Worker: worker(),
	}
}

func NewAsyncEngine(args ...string) interfaces.Engineable {
	return &AsyncEngine{
		Worker: worker(),
	}
}
