package engine

import (
	"fmt"

	. "github.com/elegardo/golden/core/models"
)

func ExampleNewAsyncEngine_gateAllOperatorEq1() {
	engine := NewAsyncEngine()
	facts := map[string]any{
		"lang":    "pt",
		"country": "br",
		"isHuman": true,
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
							Operator: EQ,
							Value:    "pt",
						},
						{
							Fact:     "country",
							Operator: EQ,
							Value:    "br",
						},
						{
							Fact:     "isHuman",
							Operator: EQ,
							Value:    true,
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

func ExampleNewAsyncEngine_gateAllOperatorEq2() {
	engine := NewAsyncEngine()
	facts := map[string]any{
		"lang":    "pt",
		"country": "br",
		"isHuman": true,
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
							Operator: EQ,
							Value:    "pt",
						},
						{
							Fact:     "isHuman",
							Operator: EQ,
							Value:    true,
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
							Fact:     "country",
							Operator: EQ,
							Value:    "br",
						},
						{
							Fact:     "isHuman",
							Operator: EQ,
							Value:    true,
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

func ExampleNewAsyncEngine_gateAllOperatorEq3() {
	engine := NewAsyncEngine()
	facts := map[string]any{
		"lang":    "pt",
		"country": "br",
		"isHuman": true,
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
							Operator: EQ,
							Value:    "es",
						},
						{
							Fact:     "country",
							Operator: EQ,
							Value:    "cl",
						},
						{
							Fact:     "isHuman",
							Operator: EQ,
							Value:    true,
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
							Fact:     "lang",
							Operator: EQ,
							Value:    "pt",
						},
						{
							Fact:     "country",
							Operator: EQ,
							Value:    "br",
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

func ExampleNewAsyncEngine_gateAllOperatorEq4() {
	engine := NewAsyncEngine()
	facts := map[string]any{
		"lang":    "pt",
		"country": "br",
		"isHuman": true,
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
							Operator: EQ,
							Value:    "pt",
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

func ExampleNewAsyncEngine_gateAllOperatorEq5() {
	engine := NewAsyncEngine()
	facts := map[string]any{
		"lang":    "pt",
		"country": "br",
		"isHuman": true,
	}
	engine.Given([]Rule{
		{
			Event: ExampleEvent{text: "Event 1"},
			Conditions: []Condition{
				{
					Gate: ALL,
					Conditionals: []Conditional{
						{
							Fact:     "isHuman",
							Operator: EQ,
							Value:    false,
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

func ExampleNewAsyncEngine_gateAllOperatorEq6() {
	engine := NewAsyncEngine()
	facts := map[string]any{
		"lang":    "pt",
		"country": "br",
		"isHuman": true,
	}
	engine.Given([]Rule{
		{
			Event: ExampleEvent{text: "Event 1"},
			Conditions: []Condition{
				{
					Gate: ALL,
					Conditionals: []Conditional{
						{
							Fact:     "other",
							Operator: EQ,
							Value:    "some",
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

func ExampleNewSyncEngine_gateAllOperatorEq1() {
	engine := NewSyncEngine()
	facts := map[string]any{
		"lang":    "pt",
		"country": "br",
		"isHuman": true,
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
							Operator: EQ,
							Value:    "pt",
						},
						{
							Fact:     "country",
							Operator: EQ,
							Value:    "br",
						},
						{
							Fact:     "isHuman",
							Operator: EQ,
							Value:    true,
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

func ExampleNewSyncEngine_gateAllOperatorEq2() {
	engine := NewSyncEngine()
	facts := map[string]any{
		"lang":    "pt",
		"country": "br",
		"isHuman": true,
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
							Operator: EQ,
							Value:    "pt",
						},
						{
							Fact:     "isHuman",
							Operator: EQ,
							Value:    true,
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
							Fact:     "country",
							Operator: EQ,
							Value:    "br",
						},
						{
							Fact:     "isHuman",
							Operator: EQ,
							Value:    true,
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

func ExampleNewSyncEngine_gateAllOperatorEq3() {
	engine := NewSyncEngine()
	facts := map[string]any{
		"lang":    "pt",
		"country": "br",
		"isHuman": true,
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
							Operator: EQ,
							Value:    "es",
						},
						{
							Fact:     "country",
							Operator: EQ,
							Value:    "cl",
						},
						{
							Fact:     "isHuman",
							Operator: EQ,
							Value:    true,
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
							Fact:     "lang",
							Operator: EQ,
							Value:    "pt",
						},
						{
							Fact:     "country",
							Operator: EQ,
							Value:    "br",
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

func ExampleNewSyncEngine_gateAllOperatorEq4() {
	engine := NewSyncEngine()
	facts := map[string]any{
		"lang":    "pt",
		"country": "br",
		"isHuman": true,
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
							Operator: EQ,
							Value:    "pt",
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

func ExampleNewSyncEngine_gateAllOperatorEq5() {
	engine := NewSyncEngine()
	facts := map[string]any{
		"lang":    "pt",
		"country": "br",
		"isHuman": true,
	}
	engine.Given([]Rule{
		{
			Event: ExampleEvent{text: "Event 1"},
			Conditions: []Condition{
				{
					Gate: ALL,
					Conditionals: []Conditional{
						{
							Fact:     "isHuman",
							Operator: EQ,
							Value:    false,
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

func ExampleNewSyncEngine_gateAllOperatorEq6() {
	engine := NewSyncEngine()
	facts := map[string]any{
		"lang":    "pt",
		"country": "br",
		"isHuman": true,
	}
	engine.Given([]Rule{
		{
			Event: ExampleEvent{text: "Event 1"},
			Conditions: []Condition{
				{
					Gate: ALL,
					Conditionals: []Conditional{
						{
							Fact:     "other",
							Operator: EQ,
							Value:    "some",
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
