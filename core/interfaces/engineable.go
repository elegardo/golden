package interfaces

import (
	"github.com/elegardo/golden/core/models"
)

// # Given(..)
//
// Es una función "Builder" que permite configurar las reglas que se usaran en la ejecucion del Engine.
//
// Puede ser usado de 2 formas:
//
//	var engine = SomeEngine()
//	engine.Given(rules)
//
// ó directamente antes del Run(..)
//
//	var engine = SomeEngine()
//	engine.Given(rules).Run(facts, callback)
//
// # Run(..)
//
// Ejecuta las reglas y acepta como parametro un map con los "facts" y una función "callback"
// que recibe elementos "Emmiter" de las reglas que son gatilladas.
//
// Puede ser usado de 2 formas:
//
//	var engine = SomeEngine()
//	engine.Given(rules).Run(facts, callback)
//
// ó si las reglas ya han sido configuradas previamente, de la forma siguiente:
//
//	var engine = SomeEngine()
//	engine.Given(rules)
//	...
//	engine.Run(facts, callback)
type Engineable interface {
	Given(rules []models.Rule) Engineable
	Run(facts map[string]any, callback models.Callback)
}
