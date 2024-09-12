package interfaces

import (
	"github.com/elegardo/golden/core/models"
)

// # Given(..)
//
// Es una función "Builder" que permite configurar las "rules" que se usaran en la ejecución del Engine.
//
//	var engine = SomeEngine()
//	engine.Given(rules).When(facts)
//
// ó en otro orden:
//
//	engine.When(facts).Given(rules)
//
// # When(..)
//
// Es una función "Builder" que permite configurar los "facts" que se usaran en la ejecucion del Engine.
//
//	var engine = SomeEngine()
//	engine.Given(rules).When(facts)
//
// ó en otro orden:
//
//	engine.When(facts).Given(rules)
//
// # Run(..)
//
// Evalua las reglas y acepta como parametro una función "callback"
// que recibe elementos "Emmiter" que son eventos de las reglas gatilladas.
//
// Para ser ejecutado se deben configurar previamente las "rules" y las "facts":
//
//	var engine = SomeEngine()
//	engine.Given(rules).When(facts).Run(callback)
//
// ó en otro orden:
//
//	engine.When(facts).Given(rules).Run(callback)
type Engine interface {
	Given(rules []models.Rule) Engine
	When(facts map[string]any) Engine
	Run(callback models.Callback)
}
