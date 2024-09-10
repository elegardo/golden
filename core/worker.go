package core

import (
	"sync"

	"github.com/elegardo/golden/core/interfaces"
	"github.com/elegardo/golden/core/models"
)

const one = 1

type Worker struct {
	Match interfaces.Matchable
}

func (w *Worker) Execute(rule *models.Rule, facts map[string]any) bool {
	group := sync.WaitGroup{}
	main := make(chan bool)

	for _, condition := range rule.Conditions {
		group.Add(one)
		go func(gate *models.Gate, conditionals *[]models.Conditional) {
			defer group.Done()
			w.evaluate(gate, main, facts, conditionals)
		}(&condition.Gate, &condition.Conditionals)
	}

	go func() {
		group.Wait()
		close(main)
	}()

	allTrue := false
	for result := range main {
		if result {
			allTrue = true
		} else {
			allTrue = false
			break
		}
	}

	return allTrue
}

func (re *Worker) evaluate(gate *models.Gate, main chan<- bool, facts map[string]any, conditionals *[]models.Conditional) {
	switch *gate {
	case models.ALL:
		re.Match.AllTrue(main, facts, conditionals)
	case models.ANY:
		re.Match.AnyTrue(main, facts, conditionals)
	case models.NONE:
		re.Match.NoneTrue(main, facts, conditionals)
	}
}
