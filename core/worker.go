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

func (w *Worker) Execute(rule models.Rule, facts map[string]any) bool {
	wg := sync.WaitGroup{}
	ch := make(chan bool, len(rule.Conditions))

	// sort.Sort(models.ConditionSorter(rule.Conditions))

	for _, condition := range rule.Conditions {
		wg.Add(one)
		go w.evaluate(&wg, ch, facts, condition)
	}

	wg.Wait()
	close(ch)

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

func (re *Worker) evaluate(wg *sync.WaitGroup, ch chan<- bool, facts map[string]any, condition models.Condition) {
	defer wg.Done()
	switch condition.Gate {
	case models.ALL:
		ch <- re.Matcher.AllTrue(facts, condition.Conditionals)
	case models.ANY:
		ch <- re.Matcher.AnyTrue(facts, condition.Conditionals)
	case models.NONE:
		ch <- re.Matcher.NoneTrue(facts, condition.Conditionals)
	}
}
