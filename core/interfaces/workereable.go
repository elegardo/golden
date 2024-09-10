package interfaces

import (
	"github.com/elegardo/golden/core/models"
)

type Workereable interface {
	Execute(rule *models.Rule, facts map[string]any) bool
}
