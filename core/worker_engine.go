package core

import (
	"github.com/elegardo/golden/core/interfaces"
	"github.com/elegardo/golden/core/domain"
)

type WorkerEngine struct {
	Matcher interfaces.Matcher
}

func (w *WorkerEngine) Execute(rule domain.Rule, facts map[string]any) bool {

	allTrue := false

	for _, condition := range rule.Conditions {
		if w.Matcher.Match(condition.Gate, facts, condition.Conditionals) {
			allTrue = true
		} else {
			allTrue = false
			break
		}
	}

	return allTrue
}
