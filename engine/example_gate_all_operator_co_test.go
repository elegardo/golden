package engine

import (
	"fmt"

	. "github.com/elegardo/golden/core/domain"
)

func ExampleNewEngine_gateAllOperatorCo1() {
	var engine = NewEngine()
	facts := map[string]any{
		"lang":    "pt",
		"country": "br",
		"points":  99,
	}
	engine.Given([]Rule{
		{
			Event: Event{text: "Event 1"},
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
	}).Run(facts, func(emmiter Emmiter) {
		fmt.Println(emmiter.Value())
	})

	// Output:
	// Event 1
}

func ExampleNewEngine_gateAllOperatorCo2() {
	var engine = NewEngine()
	facts := map[string]any{
		"lang":    "pt",
		"country": "br",
		"points":  99,
	}
	engine.Given([]Rule{
		{
			Event: Event{order: 1, text: "Event 1"},
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
			Event: Event{order: 2, text: "Event 2"},
			Conditions: []Condition{
				{
					Gate: ALL,
					Conditionals: []Conditional{
						{
							Fact:     "lang",
							Operator: CO,
							Value:    []string{"pt", "en"},
						},
						{
							Fact:     "country",
							Operator: CO,
							Value:    []string{"br", "pe"},
						},
						{
							Fact:     "points",
							Operator: CO,
							Value:    []int{9, 99, 9_999},
						},
					},
				},
			},
		},
	}).Run(facts, func(emmiter Emmiter) {
		fmt.Println(emmiter.Value())
	})

	// Unordered output:
	// Event 1
	// Event 2
}

func ExampleNewEngine_gateAllOperatorCo3() {
	var engine = NewEngine()
	facts := map[string]any{
		"lang":    "pt",
		"country": "br",
		"points":  99,
	}
	engine.Given([]Rule{
		{
			Event: Event{order: 1, text: "Event 1"},
			Conditions: []Condition{
				{
					Gate: ALL,
					Conditionals: []Conditional{
						{
							Fact:     "lang",
							Operator: CO,
							Value:    []string{"es", "en"},
						},
						{
							Fact:     "country",
							Operator: CO,
							Value:    []string{"cl", "pe"},
						},
						{
							Fact:     "points",
							Operator: CO,
							Value:    []int{1, 11, 111},
						},
					},
				},
			},
		},
		{
			Event: Event{order: 2, text: "Event 2"},
			Conditions: []Condition{
				{
					Gate: ALL,
					Conditionals: []Conditional{
						{
							Fact:     "lang",
							Operator: CO,
							Value:    []string{"pt", "en"},
						},
						{
							Fact:     "country",
							Operator: CO,
							Value:    []string{"br", "pe"},
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
	}).Run(facts, func(emmiter Emmiter) {
		fmt.Println(emmiter.Value())
	})

	// Output:
	// Event 2
}

func ExampleNewEngine_gateAllOperatorCo4() {
	var engine = NewEngine()
	facts := map[string]any{
		"lang": "pt",
	}
	engine.Given([]Rule{
		{
			Event: Event{order: 1, text: "Event 1"},
			Conditions: []Condition{
				{
					Gate: ALL,
					Conditionals: []Conditional{
						{
							Fact:     "lang",
							Operator: CO,
							Value:    []string{"pt", "en"},
						},
					},
				},
			},
		},
	}).Run(facts, func(emmiter Emmiter) {
		fmt.Println(emmiter.Value())
	})

	// Output:
	// Event 1
}

func ExampleNewEngine_gateAllOperatorCo5() {
	var engine = NewEngine()
	facts := map[string]any{
		"lang": "pt",
	}
	engine.Given([]Rule{
		{
			Event: Event{order: 1, text: "Event 1"},
			Conditions: []Condition{
				{
					Gate: ALL,
					Conditionals: []Conditional{
						{
							Fact:     "lang",
							Operator: CO,
							Value:    []string{"es", "en"},
						},
					},
				},
			},
		},
	}).Run(facts, func(emmiter Emmiter) {
		fmt.Println(emmiter.Value())
	})

	// Output:
}
