package engine

import (
	"testing"
)

func TestGate_String(t *testing.T) {
	tests := []struct {
		gate     Gate
		expected string
	}{
		{ALL, "all"},
		{ANY, "any"},
		{NONE, "none"},
	}

	for _, tt := range tests {
		result := tt.gate.String()
		if result != tt.expected {
			t.Errorf("Gate.String() = %v, want %v", result, tt.expected)
		}
	}
}

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
	}

	for _, tt := range tests {
		result := tt.operator.String()
		if result != tt.expected {
			t.Errorf("Operator.String() = %v, want %v", result, tt.expected)
		}
	}
}
