package interfaces

import "github.com/elegardo/golden/core/models"

type Evaluable interface {
	Evaluate(operator *models.Operator, fact, value any) bool
}
