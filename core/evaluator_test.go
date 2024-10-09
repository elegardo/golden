package core

import (
	"testing"

	. "github.com/elegardo/golden/core/models"
)

var mockReturn int

type mockComparable struct {
}

func (e *mockComparable) Compare(fact, value any) int {
	return mockReturn
}

func TestEvaluator_Evaluate(t *testing.T) {
	evaluator := Evaluator{
		Comparator: &mockComparable{},
	}

	tests := []struct {
		name        string
		fact        any
		conditional Conditional
		expected    bool
		mock        int
	}{
		{
			name: "Equal integers",
			fact: 10,
			conditional: Conditional{
				Operator: EQ,
				Value:    10,
			},
			expected: true,
			mock:     0,
		},
		{
			name: "Not Equal integers",
			fact: 5,
			conditional: Conditional{
				Operator: EQ,
				Value:    10,
			},
			expected: false,
			mock:     1,
		},
		{
			name: "GreaterThan integers",
			fact: 10,
			conditional: Conditional{
				Operator: GT,
				Value:    5,
			},
			expected: true,
			mock:     1,
		},
		{
			name: "Not GreaterThan integers",
			fact: 5,
			conditional: Conditional{
				Operator: GT,
				Value:    10,
			},
			expected: false,
			mock:     -1,
		},
		{
			name: "LessThan integers",
			fact: 5,
			conditional: Conditional{
				Operator: LT,
				Value:    10,
			},
			expected: true,
			mock:     -1,
		},
		{
			name: "Not LessThan integers",
			fact: 10,
			conditional: Conditional{
				Operator: LT,
				Value:    5,
			},
			expected: false,
			mock:     1,
		},
		{
			name: "Contains comparator",
			fact: "test",
			conditional: Conditional{
				Operator: IN,
				Value:    "test",
			},
			expected: true,
		},
	}

	for _, tt := range tests {
		mockReturn = tt.mock
		result := evaluator.Evaluate(tt.conditional.Operator, tt.fact, tt.conditional.Value)
		if result != tt.expected {
			t.Errorf("Evaluator.Evaluate() = %v, want %v", result, tt.expected)
		}
	}
}
