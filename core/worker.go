package core

import (
	"sync"

	"github.com/elegardo/golden/core/interfaces"
	"github.com/elegardo/golden/core/models"
)

const one = 1

type Worker struct {
	Matcher interfaces.Matchable
}

// TODO: priorizar ejecución de Gates en el siguiente orden: ALL, NONE, ANY.
// Lo anterior permite descartar mas rapido un false,
func (w *Worker) Execute(rule *models.Rule, facts map[string]any) bool {
	group := sync.WaitGroup{}
	ch := make(chan bool, len(rule.Conditions))

	for _, condition := range rule.Conditions {
		group.Add(one)
		go func(gate *models.Gate, conditionals *[]models.Conditional) {
			defer group.Done()
			ch <- w.evaluate(gate, facts, conditionals)
		}(&condition.Gate, &condition.Conditionals)
	}

	go func() {
		group.Wait()
		close(ch)
	}()

	allTrue := false
	for result := range ch {
		if result {
			allTrue = true
		} else {
			allTrue = false
			break
		}
	}

	return allTrue
}

func (re *Worker) evaluate(gate *models.Gate, facts map[string]any, conditionals *[]models.Conditional) bool {
	switch *gate {
	case models.ALL:
		return re.Matcher.AllTrue(facts, conditionals)
	case models.ANY:
		return re.Matcher.AnyTrue(facts, conditionals)
	case models.NONE:
		return re.Matcher.NoneTrue(facts, conditionals)
	default:
		return false
	}
}
