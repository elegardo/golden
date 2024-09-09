// Package con las funciones necesarias para la ejecucion de reglas
package core

import (
	"sync"

	"github.com/elegardo/golden/core/domain"
	"github.com/elegardo/golden/core/interfaces"
)

const one = 1

type value struct {
	trigger bool
	emmiter domain.Emmiter
}

// Estructura principal del motor de reglas
type Engine struct {
	Match interfaces.Matchable
	rules []domain.Rule
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

func (re *Engine) Given(rules []domain.Rule) *Engine {
	re.rules = rules
	return re
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

func (re *Engine) Run(facts map[string]any, callback domain.Callback) {
	queue := make(chan value, len(re.rules))

	for _, rule := range re.rules {
		go func(r *domain.Rule) {
			if re.worker(r, facts) {
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

func (re *Engine) worker(rule *domain.Rule, facts map[string]any) bool {
	group := sync.WaitGroup{}
	main := make(chan bool)

	for _, condition := range rule.Conditions {
		group.Add(one)
		go func(gate *domain.Gate, conditionals *[]domain.Conditional) {
			defer group.Done()
			re.evaluate(gate, main, facts, conditionals)
		}(&condition.Gate, &condition.Conditionals)
	}

	go func() {
		group.Wait()
		close(main)
	}()

	allTrue := false
	for result := range main {
		if result {
			allTrue = true
		} else {
			allTrue = false
			break
		}
	}

	return allTrue
}

func (re *Engine) evaluate(gate *domain.Gate, main chan<- bool, facts map[string]any, conditionals *[]domain.Conditional) {
	switch *gate {
	case domain.ALL:
		re.Match.AllTrue(main, facts, conditionals)
	case domain.ANY:
		re.Match.AnyTrue(main, facts, conditionals)
	case domain.NONE:
		re.Match.NoneTrue(main, facts, conditionals)
	}
}
