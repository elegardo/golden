package interfaces

import "github.com/elegardo/golden/core/models"

type Runnable interface {
	Run(rule *models.Rule, facts map[string]any) bool
}
