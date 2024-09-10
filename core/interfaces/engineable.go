package interfaces

import (
	"github.com/elegardo/golden/core/models"
)

type Engineable interface {
	Given(rules []models.Rule) Engineable
	Run(facts map[string]any, callback models.Callback)
}
