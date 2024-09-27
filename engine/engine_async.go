package engine

import (
	"sync"

	"github.com/elegardo/golden/core/interfaces"
	"github.com/elegardo/golden/core/models"
)

// Create a wait group to synchronize workers
var wg sync.WaitGroup

const numWorkers = 5

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

func (re *AsyncEngine) Run(callback models.Callback) {
	numJobs := len(re.rules)

	// Create channels for jobs and results
	jobs := make(chan models.Rule, numJobs)
	results := make(chan models.Emmiter, numJobs)

	// Launch worker goroutines
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		//Worker
		go func() {
			defer wg.Done()

			for rule := range jobs {
				if re.Worker.Execute(rule, re.facts) {
					results <- rule.Event
				}
			}

		}()
	}

	// Generate jobs
	for _, rule := range re.rules {
		jobs <- rule
	}

	// Close the jobs channel to signal that no more tasks will be added
	close(jobs)

	// Wait for all workers to finish
	wg.Wait()

	// Close the results channel
	close(results)

	for result := range results {
		callback(result)
	}
}
