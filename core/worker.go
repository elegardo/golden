package core

import (
	"github.com/elegardo/golden/core/interfaces"
	"github.com/elegardo/golden/core/models"
)

type Worker struct {
	Matcher interfaces.Matchable
}

func (w *Worker) Execute(rule models.Rule, facts map[string]any) bool {

	allTrue := false

	for _, condition := range rule.Conditions {
		if w.Matcher.Execute(condition.Gate, facts, condition.Conditionals) {
			allTrue = true
		} else {
			allTrue = false
			break
		}
	}

	return allTrue
}
