package core

import (
	"testing"
)

func TestContains(t *testing.T) {
	container := FactContainer[string]{}

	tests := []struct {
		name     string
		fact     string
		value    []string
		expected bool
	}{
		{
			name:     "Contains true array string",
			fact:     "apple",
			value:    []string{"banana", "apple"},
			expected: true,
		},
	}

	for _, tt := range tests {
		result := container.Contains(tt.fact, tt.value)
		if result != tt.expected {
			t.Errorf("Comparator.Contains() = %v, want %v", result, tt.expected)
		}
	}
}
