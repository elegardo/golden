package engine

import (
	"sync"

	"github.com/elegardo/golden/core/interfaces"
	"github.com/elegardo/golden/core/models"
)

type AsyncEngine struct {
	Worker interfaces.Workereable
	rules  []models.Rule
	facts  map[string]any
}

func (re *AsyncEngine) Given(rules []models.Rule) interfaces.Engine {
	re.rules = rules
	return re
}

func (re *AsyncEngine) When(facts map[string]any) interfaces.Engine {
	re.facts = facts
	return re
}

// 6 allocs/op
func (re *AsyncEngine) Run(callback models.Callback) {
	wg := sync.WaitGroup{}
	ch := make(chan models.Emmiter, len(re.rules))

	for _, rule := range re.rules {
		wg.Add(1)
		go re.worker(&wg, ch, rule)
	}

	go func() {
		wg.Wait()
		defer close(ch)
	}()

	for result := range ch {
		callback(result)
	}
}

func (re *AsyncEngine) worker(wg *sync.WaitGroup, ch chan<- models.Emmiter, rule models.Rule) {
	defer wg.Done()
	if re.Worker.Execute(rule, re.facts) {
		ch <- rule.Event
	}
}
