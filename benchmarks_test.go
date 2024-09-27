package engine

import (
	"testing"

	. "github.com/elegardo/golden/core/models"
	. "github.com/elegardo/golden/engine"
)

type event struct {
	order int
	text  any
}

func (e event) Value() any {
	return e.text
}

func (e event) Order() int {
	return e.order
}

func makeRules() []Rule {
	var rules = []Rule{
		{
			Id:    1,
			Name:  "Name is John and Age is greater than 25",
			Event: event{text: "Event 1"},
			Conditions: []Condition{
				{
					Gate: ALL,
					Conditionals: []Conditional{
						{
							Fact:     "age",
							Operator: GT,
							Value:    20,
						},
						{
							Fact:     "name",
							Operator: EQ,
							Value:    "John",
						},
					},
				},
				{
					Gate: ANY,
					Conditionals: []Conditional{
						{
							Fact:     "country",
							Operator: EQ,
							Value:    "cl",
						},
						{
							Fact:     "country",
							Operator: EQ,
							Value:    "pe",
						},
					},
				},
				{
					Gate: NONE,
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
		{
			Id:    2,
			Name:  "Name is John and Age is greater than 25",
			Event: event{text: "Event 2"},
			Conditions: []Condition{
				{
					Gate: ALL,
					Conditionals: []Conditional{
						{
							Fact:     "age",
							Operator: GT,
							Value:    20,
						},
						{
							Fact:     "name",
							Operator: EQ,
							Value:    "John",
						},
						{
							Fact:     "age",
							Operator: CO,
							Value:    []int{20, 30},
						},
					},
				},
				{
					Gate: ANY,
					Conditionals: []Conditional{
						{
							Fact:     "country",
							Operator: CO,
							Value:    []string{"pe", "cl"},
						},
					},
				},
				{
					Gate: NONE,
					Conditionals: []Conditional{
						{
							Fact:     "country",
							Operator: IN,
							Value:    "peru",
						},
					},
				},
			},
		},
	}

	var result []Rule

	for i := 0; i < 5000; i++ {
		result = append(result, rules[0])
		result = append(result, rules[1])
	}

	return result
}

func BenchmarkAsyncEngine(b *testing.B) {
	var rules = makeRules()
	var engine = NewAsyncEngine()
	facts := map[string]any{
		"age":     30,
		"name":    "John",
		"country": "cl",
		"lang":    "es",
	}
	engine.Given(rules).When(facts)
	for i := 0; i < b.N; i++ {
		engine.Run(func(emmiter Emmiter) {})
	}
}

func BenchmarkSyncEngine(b *testing.B) {
	var rules = makeRules()
	var engine = NewSyncEngine()
	facts := map[string]any{
		"age":     30,
		"name":    "John",
		"country": "cl",
		"lang":    "es",
	}
	engine.Given(rules).When(facts)
	for i := 0; i < b.N; i++ {
		engine.Run(func(emmiter Emmiter) {})
	}
}
