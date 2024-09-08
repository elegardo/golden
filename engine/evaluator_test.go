package engine

import (
	"testing"
)

func TestEvaluator_Evaluate(t *testing.T) {
	evaluator := Evaluator{}

	tests := []struct {
		name        string
		fact        any
		conditional Conditional
		expected    bool
	}{
		{
			name: "Equal integers",
			fact: 10,
			conditional: Conditional{
				Operator: EQ,
				Value:    10,
			},
			expected: true,
		},
		{
			name: "Not Equal integers",
			fact: 5,
			conditional: Conditional{
				Operator: EQ,
				Value:    10,
			},
			expected: false,
		},
		{
			name: "GreaterThan integers",
			fact: 10,
			conditional: Conditional{
				Operator: GT,
				Value:    5,
			},
			expected: true,
		},
		{
			name: "Not GreaterThan integers",
			fact: 5,
			conditional: Conditional{
				Operator: GT,
				Value:    10,
			},
			expected: false,
		},
		{
			name: "LessThan integers",
			fact: 5,
			conditional: Conditional{
				Operator: LT,
				Value:    10,
			},
			expected: true,
		},
		{
			name: "Not LessThan integers",
			fact: 10,
			conditional: Conditional{
				Operator: LT,
				Value:    5,
			},
			expected: false,
		},
		{
			name: "Contains comparator",
			fact: "test",
			conditional: Conditional{
				Operator: CO,
				Value:    "test",
			},
			expected: true,
		},
	}

	for _, tt := range tests {
		result := evaluator.Evaluate(&tt.conditional.Operator, tt.fact, tt.conditional.Value)
		if result != tt.expected {
			t.Errorf("Evaluator.Evaluate() = %v, want %v", result, tt.expected)
		}
	}
}
