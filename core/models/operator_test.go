package models

import (
	"testing"
)

func TestOperator_String(t *testing.T) {
	tests := []struct {
		operator Operator
		expected string
	}{
		{EQ, "eq"},
		{NE, "ne"},
		{GT, "gt"},
		{GE, "ge"},
		{LT, "lt"},
		{LE, "le"},
		{CO, "co"},
		{NC, "nc"},
		{IN, "in"},
		{NI, "ni"},
	}

	for _, tt := range tests {
		result := tt.operator.String()
		if result != tt.expected {
			t.Errorf("Operator.String() = %v, want %v", result, tt.expected)
		}
	}
}
