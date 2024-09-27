package engine

import (
	"github.com/elegardo/golden/core"
	"github.com/elegardo/golden/core/interfaces"
)

// TODO: implement CLI flags
// -workers-pool=5
// -exit-first-success
// -exit-first-fail

func syncWorker() interfaces.Workereable {
	return &core.Worker{
		Matcher: &core.Matcher{
			Evaluator: &core.Evaluator{},
		},
	}
}

func NewRunnerEngine(args ...string) interfaces.Runnable {
	return &RunnerEngine{
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
