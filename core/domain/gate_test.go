package domain

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
