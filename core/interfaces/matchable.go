package interfaces

import "github.com/elegardo/golden/core/models"

// Declara la funcionalidad para la busqueda por Gate: AllTrue, AnyTrue y NoneTrue
type Matchable interface {
	AllTrue(main chan<- bool, facts map[string]any, conditionals *[]models.Conditional)
	AnyTrue(main chan<- bool, facts map[string]any, conditionals *[]models.Conditional)
	NoneTrue(main chan<- bool, facts map[string]any, conditionals *[]models.Conditional)
}
