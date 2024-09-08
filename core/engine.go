// Package engine provides structs, functions and interfaces for rule engine
package core

import (
	"sync"
)

const one = 1

type emmit struct {
	trigger bool
	event   Emmiter
}

type Engine struct {
	Match Matchable
	rules []Rule
}

func (re *Engine) Given(rules []Rule) *Engine {
	re.rules = rules
	return re
}

func (re *Engine) Run(facts map[string]any, callback Callback) {
	queue := make(chan emmit, len(re.rules))

	for _, rule := range re.rules {
		go func(r *Rule) {
			if re.worker(r, facts) {
				queue <- emmit{trigger: true, event: r.Event}
			} else {
				queue <- emmit{trigger: false, event: r.Event}
			}
		}(&rule)
	}

	for i := 0; i < cap(queue); i++ {
		emmit := <-queue
		if emmit.trigger {
			callback(emmit.event)
		}
	}

}

func (re *Engine) worker(rule *Rule, facts map[string]any) bool {
	group := sync.WaitGroup{}
	main := make(chan bool)

	for _, condition := range rule.Conditions {
		group.Add(one)
		go func(gate *Gate, conditionals *[]Conditional) {
			defer group.Done()
			re.evaluate(gate, main, facts, conditionals)
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

func (re *Engine) evaluate(gate *Gate, main chan<- bool, facts map[string]any, conditionals *[]Conditional) {
	switch *gate {
	case ALL:
		re.Match.allTrue(main, facts, conditionals)
	case ANY:
		re.Match.anyTrue(main, facts, conditionals)
	case NONE:
		re.Match.noneTrue(main, facts, conditionals)
	}
}
