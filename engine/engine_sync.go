// Package con las funciones necesarias para la ejecucion de reglas
package engine

import (
	"github.com/elegardo/golden/core/interfaces"
	"github.com/elegardo/golden/core/models"
)

type SyncEngine struct {
	Worker interfaces.Workereable
	rules  []models.Rule
}

func (se *SyncEngine) Given(rules []models.Rule) interfaces.Engineable {
	se.rules = rules
	return se
}

func (se *SyncEngine) Run(facts map[string]any, callback models.Callback) {
	//TODO: sort by priority

	for _, rule := range se.rules {
		if se.Worker.Execute(&rule, facts) {
			callback(rule.Event)
		}
	}

}
