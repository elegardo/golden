// Package con las funciones necesarias para la ejecucion de reglas
package core

import (
	"github.com/elegardo/golden/core/interfaces"
	"github.com/elegardo/golden/core/domain"
)

type RunnerEngine struct {
	Worker interfaces.Worker
}

func (se *RunnerEngine) Run(rule domain.Rule, facts map[string]any) bool {
	return se.Worker.Execute(rule, facts)
}
