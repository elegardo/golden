package engine

import (
	"fmt"

	. "github.com/elegardo/golden/core/models"
)

var rules2 = []Rule{
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
}

func ExampleNewAsyncEngine_gateAllOperatorEq1() {
	engine := NewAsyncEngine()
	facts := map[string]any{
		"lang":    "pt",
		"country": "br",
		"isHuman": true,
	}
	engine.Given(rules2).When(facts).Run(func(emmiter Emmiter) {
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
	}).When(facts).Run(func(emmiter Emmiter) {
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
	}).When(facts).Run(func(emmiter Emmiter) {
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
	}).When(facts).Run(func(emmiter Emmiter) {
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
	}).When(facts).Run(func(emmiter Emmiter) {
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
	}).When(facts).Run(func(emmiter Emmiter) {
		fmt.Println(emmiter.Value())
	})

	// Output:
}
