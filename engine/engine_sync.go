package engine

import (
	"github.com/elegardo/golden/core/interfaces"
	"github.com/elegardo/golden/core/models"
)

type SyncEngine struct {
	Worker interfaces.Workereable
	rules  []models.Rule
	facts  map[string]any
}

func (re *SyncEngine) Given(rules []models.Rule) interfaces.Engine {
	re.rules = rules
	return re
}

func (re *SyncEngine) When(facts map[string]any) interfaces.Engine {
	re.facts = facts
	return re
}

func (re *SyncEngine) Run(callback models.Callback) {
	for _, rule := range re.rules {
		if re.Worker.Execute(rule, re.facts) {
			callback(rule.Event)
		}
	}
}
