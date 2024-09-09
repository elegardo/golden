package core

import (
	"fmt"
	"reflect"
	"testing"

	. "github.com/elegardo/golden/core/domain"
)

type MockEvaluator struct {
	MockResults map[string]bool
}

func formatKey(x, y any) string {
	return fmt.Sprintf("%v:%v", x, y)
}

func (me *MockEvaluator) Evaluate(operator *Operator, x, y any) bool {
	key := formatKey(x, y)
	if result, exists := me.MockResults[key]; exists {
		return result
	}
	return false
}

func TestRuleEvaluator_Execute(t *testing.T) {
	tests := []struct {
		name         string
		facts        map[string]any
		conditionals []Conditional
		mockResults  map[string]bool
		expected     []bool
	}{
		{
			name: "Single fact matches condition",
			facts: map[string]any{
				"age": 30,
			},
			conditionals: []Conditional{
				{Fact: "age", Operator: GT, Value: 25},
			},
			mockResults: map[string]bool{
				"30:25": true,
			},
			expected: []bool{true},
		},
		{
			name: "Single fact does not match condition",
			facts: map[string]any{
				"age": 20,
			},
			conditionals: []Conditional{
				{Fact: "age", Operator: GT, Value: 25},
			},
			mockResults: map[string]bool{
				"20:25": false,
			},
			expected: []bool{false},
		},
		{
			name: "Fact not found",
			facts: map[string]any{
				"name": "John",
			},
			conditionals: []Conditional{
				{Fact: "age", Operator: GT, Value: 25},
			},
			mockResults: map[string]bool{},
			expected:    []bool{},
		},
		{
			name: "Multiple conditionals, all match",
			facts: map[string]any{
				"age":  30,
				"name": "John",
			},
			conditionals: []Conditional{
				{Fact: "age", Operator: GT, Value: 25},
				{Fact: "name", Operator: EQ, Value: "John"},
			},
			mockResults: map[string]bool{
				"30:25":     true,
				"John:John": true,
			},
			expected: []bool{true, true},
		},
		{
			name: "Multiple conditionals, some do not match",
			facts: map[string]any{
				"age":  20,
				"name": "John",
			},
			conditionals: []Conditional{
				{Fact: "age", Operator: GT, Value: 25},
				{Fact: "name", Operator: EQ, Value: "John"},
			},
			mockResults: map[string]bool{
				"20:25":     false,
				"John:John": true,
			},
			expected: []bool{false, true},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockEvaluator := &MockEvaluator{
				MockResults: tt.mockResults,
			}

			ruleEvaluator := RuleEvaluator{Evaluator: mockEvaluator}

			ch := make(chan bool)
			go ruleEvaluator.Execute(ch, tt.facts, &tt.conditionals)

			var results []bool
			for result := range ch {
				results = append(results, result)
			}

			if results == nil {
				results = []bool{}
			}

			if !reflect.DeepEqual(results, tt.expected) {
				t.Errorf("RuleEvaluator.Execute() = %v, want %v", results, tt.expected)
			}
		})
	}
}
