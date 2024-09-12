package interfaces

import "github.com/elegardo/golden/core/models"

// Declara la funcionalidad para la busqueda por Gate: AllTrue, AnyTrue y NoneTrue
type Matchable interface {
	AllTrue(facts map[string]any, conditionals *[]models.Conditional) bool
	AnyTrue(facts map[string]any, conditionals *[]models.Conditional) bool
	NoneTrue(facts map[string]any, conditionals *[]models.Conditional) bool
}
