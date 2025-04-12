package interfaces

import "github.com/elegardo/golden/core/domain"

type Runner interface {
	Run(rule domain.Rule, facts map[string]any) bool
}
