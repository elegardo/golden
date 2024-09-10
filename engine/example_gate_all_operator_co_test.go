package engine

import (
	"fmt"

	. "github.com/elegardo/golden/core/models"
)

func ExampleNewAsyncEngine_gateAllOperatorCo1() {
	var engine = NewAsyncEngine()
	facts := map[string]any{
		"lang":    "pt",
		"country": "br",
		"points":  99,
	}
	engine.Given([]Rule{
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
	}).Run(facts, func(emmiter Emmiter) {
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
	}
	engine.Given([]Rule{
		{
			Event: ExampleEvent{order: 1, text: "Event 1"},
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
			Event: ExampleEvent{order: 2, text: "Event 2"},
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

func ExampleNewAsyncEngine_gateAllOperatorCo3() {
	var engine = NewAsyncEngine()
	facts := map[string]any{
		"lang":    "pt",
		"country": "br",
		"points":  99,
	}
	engine.Given([]Rule{
		{
			Event: ExampleEvent{order: 1, text: "Event 1"},
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
			Event: ExampleEvent{order: 2, text: "Event 2"},
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

func ExampleNewAsyncEngine_gateAllOperatorCo4() {
	var engine = NewAsyncEngine()
	facts := map[string]any{
		"lang": "pt",
	}
	engine.Given([]Rule{
		{
			Event: ExampleEvent{order: 1, text: "Event 1"},
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

func ExampleNewAsyncEngine_gateAllOperatorCo5() {
	var engine = NewAsyncEngine()
	facts := map[string]any{
		"lang": "pt",
	}
	engine.Given([]Rule{
		{
			Event: ExampleEvent{order: 1, text: "Event 1"},
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

// SYNC ENGINE

func ExampleNewSyncEngine_gateAllOperatorCo1() {
	var engine = NewSyncEngine()
	facts := map[string]any{
		"lang":    "pt",
		"country": "br",
		"points":  99,
	}
	engine.Given([]Rule{
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
	}).Run(facts, func(emmiter Emmiter) {
		fmt.Println(emmiter.Value())
	})

	// Output:
	// Event 1
}

func ExampleNewSyncEngine_gateAllOperatorCo2() {
	var engine = NewSyncEngine()
	facts := map[string]any{
		"lang":    "pt",
		"country": "br",
		"points":  99,
	}
	engine.Given([]Rule{
		{
			Event: ExampleEvent{order: 1, text: "Event 1"},
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
			Event: ExampleEvent{order: 2, text: "Event 2"},
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

	// Output:
	// Event 1
	// Event 2
}

func ExampleNewSyncEngine_gateAllOperatorCo3() {
	var engine = NewSyncEngine()
	facts := map[string]any{
		"lang":    "pt",
		"country": "br",
		"points":  99,
	}
	engine.Given([]Rule{
		{
			Event: ExampleEvent{order: 1, text: "Event 1"},
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
			Event: ExampleEvent{order: 2, text: "Event 2"},
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

func ExampleNewSyncEngine_gateAllOperatorCo4() {
	var engine = NewSyncEngine()
	facts := map[string]any{
		"lang": "pt",
	}
	engine.Given([]Rule{
		{
			Event: ExampleEvent{order: 1, text: "Event 1"},
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

func ExampleNewSyncEngine_gateAllOperatorCo5() {
	var engine = NewSyncEngine()
	facts := map[string]any{
		"lang": "pt",
	}
	engine.Given([]Rule{
		{
			Event: ExampleEvent{order: 1, text: "Event 1"},
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
