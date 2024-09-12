package engine

import (
	"fmt"

	. "github.com/elegardo/golden/core/models"
)

var rules1 = []Rule{
	{
		Event: ExampleEvent{text: "Event 1"},
		Conditions: []Condition{
			{
				Gate: ALL,
				Conditionals: []Conditional{
					{
						Fact:     "lang",
						Operator: CO,
						Value:    []string{"pt", "es"},
					},
					{
						Fact:     "country",
						Operator: CO,
						Value:    []string{"br", "cl"},
					},
					{
						Fact:     "points",
						Operator: CO,
						Value:    []int{9, 99, 999},
					},
				},
			},
		},
	},
	{
		Event: ExampleEvent{text: "Event 2"},
		Conditions: []Condition{
			{
				Gate: ALL,
				Conditionals: []Conditional{
					{
						Fact:     "type",
						Operator: CO,
						Value:    []string{"basic", "medium"},
					},
				},
			},
		},
	},
}

// AsyncEngine
func ExampleNewAsyncEngine_gateAllOperatorCo1() {
	engine := NewAsyncEngine()
	facts := map[string]any{
		"lang":    "pt",
		"country": "br",
		"points":  99,
	}

	engine.Given(rules1).When(facts).Run(func(emmiter Emmiter) {
		fmt.Println(emmiter.Value())
	})

	// Output:
	// Event 1
}

func ExampleNewAsyncEngine_gateAllOperatorCo2() {
	var engine = NewAsyncEngine()
	facts := map[string]any{
		"lang":    "pt",
		"country": "br",
		"points":  99,
		"type":    "basic",
	}

	engine.Given(rules1).When(facts).Run(func(emmiter Emmiter) {
		fmt.Println(emmiter.Value())
	})

	// Unordered output:
	// Event 1
	// Event 2
}

func ExampleNewAsyncEngine_gateAllOperatorCo3() {
	var engine = NewAsyncEngine()
	facts := map[string]any{
		"lang":    "en",
		"country": "cl",
		"points":  100,
		"type":    "medium",
	}

	engine.Given(rules1).When(facts).Run(func(emmiter Emmiter) {
		fmt.Println(emmiter.Value())
	})

	// Output:
	// Event 2
}

func ExampleNewAsyncEngine_gateAllOperatorCo4() {
	var engine = NewAsyncEngine()
	facts := map[string]any{
		"lang": "en",
	}

	engine.Given(rules1).When(facts).Run(func(emmiter Emmiter) {
		fmt.Println(emmiter.Value())
	})

	// Output:
}

func ExampleNewAsyncEngine_gateAllOperatorCo5() {
	var engine = NewAsyncEngine()
	facts := map[string]any{
		"type": "premium",
	}

	engine.Given(rules1).When(facts).Run(func(emmiter Emmiter) {
		fmt.Println(emmiter.Value())
	})

	// Output:
}

// RunnerEngine
func ExampleNewRunnerEngine_gateAllOperatorCo1() {
	runner := NewRunnerEngine()
	facts := map[string]any{
		"lang":    "pt",
		"country": "br",
		"points":  99,
	}

	for _, rule := range rules1 {
		fmt.Println(runner.Run(&rule, facts))
	}

	// Output:
	// true
	// false
}

func ExampleNewRunnerEngine_gateAllOperatorCo2() {
	runner := NewRunnerEngine()
	facts := map[string]any{
		"lang":    "pt",
		"country": "br",
		"points":  99,
		"type":    "basic",
	}

	for _, rule := range rules1 {
		fmt.Println(runner.Run(&rule, facts))
	}

	// Output:
	// true
	// true
}

func ExampleNewRunnerEngine_gateAllOperatorCo3() {
	runner := NewRunnerEngine()
	facts := map[string]any{
		"lang":    "en",
		"country": "cl",
		"points":  100,
		"type":    "medium",
	}

	for _, rule := range rules1 {
		fmt.Println(runner.Run(&rule, facts))
	}

	// Output:
	// false
	// true
}

func ExampleNewRunnerEngine_gateAllOperatorCo4() {
	runner := NewRunnerEngine()
	facts := map[string]any{
		"lang": "en",
	}

	for _, rule := range rules1 {
		fmt.Println(runner.Run(&rule, facts))
	}

	// Output:
	// false
	// false
}

func ExampleNewRunnerEngine_gateAllOperatorCo5() {
	runner := NewRunnerEngine()
	facts := map[string]any{
		"type": "premium",
	}

	for _, rule := range rules1 {
		fmt.Println(runner.Run(&rule, facts))
	}

	// Output:
	// false
	// false
}
