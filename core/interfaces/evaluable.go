package interfaces

import "github.com/elegardo/golden/core/domain"

type Evaluable interface {
	Evaluate(operator *domain.Operator, fact, value any) bool
}
