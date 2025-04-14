package core

import (
	"testing"

	. "github.com/elegardo/golden/core/domain"
)

var mockReturnInteger int
var mockReturnBool bool

type mockComparable struct {
}

func (e *mockComparable) Compare(pair *Pair) int {
	return mockReturnInteger
}

func (e *mockComparable) Contains(pair *Pair) bool {
	return mockReturnBool
}

func TestEvaluator_Evaluate(t *testing.T) {
	evaluator := RuleEvaluator{
		Comparator: &mockComparable{},
	}

	tests := []struct {
		name        string
		fact        any
		conditional Conditional
		expected    bool
		mockInt     int
	}{
		{
			name: "Equal integers",
			fact: 10,
			conditional: Conditional{
				Operator: EQ,
				Value:    10,
			},
			expected: true,
			mockInt:  0,
		},
		{
			name: "Not Equal integers",
			fact: 5,
			conditional: Conditional{
				Operator: EQ,
				Value:    10,
			},
			expected: false,
			mockInt:  1,
		},
		{
			name: "GreaterThan integers",
			fact: 10,
			conditional: Conditional{
				Operator: GT,
				Value:    5,
			},
			expected: true,
			mockInt:  1,
		},
		{
			name: "Not GreaterThan integers",
			fact: 5,
			conditional: Conditional{
				Operator: GT,
				Value:    10,
			},
			expected: false,
			mockInt:  -1,
		},
		{
			name: "LessThan integers",
			fact: 5,
			conditional: Conditional{
				Operator: LT,
				Value:    10,
			},
			expected: true,
			mockInt:  -1,
		},
		{
			name: "Not LessThan integers",
			fact: 10,
			conditional: Conditional{
				Operator: LT,
				Value:    5,
			},
			expected: false,
			mockInt:  1,
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
		mockReturnInteger = tt.mockInt
		pair := NewPair(tt.fact, tt.conditional.Value)
		result := evaluator.Evaluate(tt.conditional.Operator, pair)
		if result != tt.expected {
			t.Errorf("Evaluator.Evaluate() = %v, want %v", result, tt.expected)
		}
	}
}
