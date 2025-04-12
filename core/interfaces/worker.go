package interfaces

import (
	"github.com/elegardo/golden/core/domain"
)

type Worker interface {
	Execute(rule domain.Rule, facts map[string]any) bool
}
