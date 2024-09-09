package interfaces

import "github.com/elegardo/golden/core/domain"

// Declara la funcionalidad para la busqueda por Gate: AllTrue, AnyTrue y NoneTrue
type Matchable interface {
	AllTrue(main chan<- bool, facts map[string]any, conditionals *[]domain.Conditional)
	AnyTrue(main chan<- bool, facts map[string]any, conditionals *[]domain.Conditional)
	NoneTrue(main chan<- bool, facts map[string]any, conditionals *[]domain.Conditional)
}
