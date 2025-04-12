package engine

import (
	"github.com/elegardo/golden/core"
	"github.com/elegardo/golden/core/interfaces"
)

// TODO: implement CLI flags
// -workers-pool=5
// -exit-first-success
// -exit-first-fail

func syncWorker() interfaces.Worker {
	return &core.WorkerEngine{
		Matcher: &core.GateMatcher{
			Evaluator: &core.RuleEvaluator{
				Comparator: &core.FactComparator{},
			},
		},
	}
}

func NewRunnerEngine(args ...string) interfaces.Runner {
	return &core.RunnerEngine{
		Worker: syncWorker(),
	}
}

func NewAsyncEngine(args ...string) interfaces.Engine {
	return &AsyncEngine{
		Worker:  syncWorker(),
		workers: 5,
	}
}

func NewSyncEngine(args ...string) interfaces.Engine {
	return &SyncEngine{
		Worker: syncWorker(),
	}
}
