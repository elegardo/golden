// Package con las funciones necesarias para la ejecucion de reglas
package engine

import (
	"github.com/elegardo/golden/core/interfaces"
	"github.com/elegardo/golden/core/models"
)

type syncValue struct {
	trigger bool
	emmiter models.Emmiter
}

// Estructura principal del motor de reglas
type SyncEngine struct {
	Worker interfaces.Workereable
	rules  []models.Rule
}

// Given es una función "Builder" permite configurar las reglas que se usaran en la ejecucion del Engine.
//
// Puede ser usado de 2 formas:
//
//	var engine = NewEngine()
//	engine.Given(rules)
//
// ó directamente antes del Run(..)
//
//	var engine = NewEngine()
//	engine.Given(rules).Run(facts, callback)
func (se *SyncEngine) Given(rules []models.Rule) interfaces.Engineable {
	se.rules = rules
	return se
}

// Run ejecuta las reglas ya configuradas y acepta como parametro un map con los "facts" y una función "callback" que recibira elementos "Emmiter" asincronicamente.
//
// Puede ser usado de 2 formas:
//
//	var engine = NewEngine()
//	engine.Given(rules).Run(facts, callback)
//
// ó si ya han sido configuradas las reglas anteriormente, de la forma siguiente:
//
//	var engine = NewEngine()
//	engine.Given(rules)
//	...
//	engine.Run(facts, callback)
func (se *SyncEngine) Run(facts map[string]any, callback models.Callback) {
	queue := make(chan syncValue, len(se.rules))

	//TODO: sort by priority

	for _, rule := range se.rules {
		if se.Worker.Execute(&rule, facts) {
			queue <- syncValue{trigger: true, emmiter: rule.Event}
		} else {
			queue <- syncValue{trigger: false, emmiter: rule.Event}
		}
	}

	for i := 0; i < cap(queue); i++ {
		value := <-queue
		if value.trigger {
			callback(value.emmiter)
		}
	}
}
