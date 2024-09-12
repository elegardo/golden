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
				},
			},
			{
				Gate: ANY,
				Conditionals: []Conditional{
					{
						Fact:     "country",
						Operator: EQ,
						Value:    "pe",
					},
					{
						Fact:     "country",
						Operator: EQ,
						Value:    "cl",
					},
				},
			},
		},
	},
}

func BenchmarkRuleEngine(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// Inicializar el RuleEngine
		var ruleEngine = NewAsyncEngine()
		facts := map[string]any{
			"age":     30,
			"name":    "John",
			"country": "cl",
			"lang":    "es",
		}
		ruleEngine.Given(rules).When(facts).Run(some)
	}
}

func some(event Emmiter) {
}
