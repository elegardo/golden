package engine

import (
	"github.com/elegardo/golden/core/interfaces"
	"github.com/elegardo/golden/core/domain"
)

type SyncEngine struct {
	Worker interfaces.Worker
	rules  []domain.Rule
	facts  map[string]any
}

func (re *SyncEngine) Given(rules []domain.Rule) interfaces.Engine {
	re.rules = rules
	return re
}

func (re *SyncEngine) When(facts map[string]any) interfaces.Engine {
	re.facts = facts
	return re
}

func (re *SyncEngine) Run(callback domain.Callback) {
	for _, rule := range re.rules {
		if re.Worker.Execute(rule, re.facts) {
			callback(rule.Event)
		}
	}
}
