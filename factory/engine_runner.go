// Package con las funciones necesarias para la ejecucion de reglas
package engine

import (
	"github.com/elegardo/golden/core/interfaces"
	"github.com/elegardo/golden/core/models"
)

type RunnerEngine struct {
	Worker interfaces.Workereable
}

func (se *RunnerEngine) Run(rule models.Rule, facts map[string]any) bool {
	return se.Worker.Execute(rule, facts)
}
