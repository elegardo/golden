package engine

import (
	"github.com/elegardo/golden/core/interfaces"
	"github.com/elegardo/golden/core/models"
)

type asyncValue struct {
	trigger bool
	emmiter models.Emmiter
}

type AsyncEngine struct {
	Worker interfaces.Workereable
	rules  []models.Rule
}

func (re *AsyncEngine) Given(rules []models.Rule) interfaces.Engineable {
	re.rules = rules
	return re
}

func (re *AsyncEngine) Run(facts map[string]any, callback models.Callback) {
	queue := make(chan asyncValue, len(re.rules))

	for _, rule := range re.rules {
		go func(r *models.Rule) {
			if re.Worker.Execute(r, facts) {
				queue <- asyncValue{trigger: true, emmiter: r.Event}
			} else {
				queue <- asyncValue{trigger: false, emmiter: r.Event}
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
