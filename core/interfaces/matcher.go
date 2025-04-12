package interfaces

import "github.com/elegardo/golden/core/domain"

// Declara la funcionalidad para la busqueda por Gate: AllTrue, AnyTrue y NoneTrue
type Matcher interface {
	Match(gate domain.Gate, facts map[string]any, conditionals []domain.Conditional) bool
}
