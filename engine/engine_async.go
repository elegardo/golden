package engine

import (
	"github.com/elegardo/golden/core/interfaces"
	"github.com/elegardo/golden/core/models"
)

type value struct {
	trigger bool
	emmiter models.Emmiter
}

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
	queue := make(chan value, len(re.rules))

	for _, rule := range re.rules {
		go func(r *models.Rule) {
			if re.Worker.Execute(r, re.facts) {
				queue <- value{trigger: true, emmiter: r.Event}
			} else {
				queue <- value{trigger: false, emmiter: r.Event}
			}
		}(&rule)
	}

	for i := 0; i < cap(queue); i++ {
		value := <-queue
		if value.trigger {
			callback(value.emmiter)
		}
	}
}
