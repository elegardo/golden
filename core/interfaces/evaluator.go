package interfaces

import "github.com/elegardo/golden/core/domain"

type Evaluator interface {
	Evaluate(operator domain.Operator, factValue, conditionalValue any) bool
}
