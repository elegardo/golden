package core

import (
	"testing"
)

func TestCompare(t *testing.T) {
	tests := []struct {
		name     string
		fact     any
		value    any
		expected int
	}{
		{
			name:     "Compare equal integers",
			fact:     10,
			value:    10,
			expected: 0,
		},
		{
			name:     "Compare integers a > b",
			fact:     10,
			value:    5,
			expected: 1,
		},
		{
			name:     "Compare integers a < b",
			fact:     5,
			value:    10,
			expected: -1,
		},
		{
			name:     "Compare equal strings",
			fact:     "apple",
			value:    "apple",
			expected: 0,
		},
		{
			name:     "Compare strings a > b",
			fact:     "banana",
			value:    "apple",
			expected: 1,
		},
		{
			name:     "Compare strings a < b",
			fact:     "apple",
			value:    "banana",
			expected: -1,
		},
	}

	for _, tt := range tests {
		result := Compare(tt.fact, tt.value)
		if result != tt.expected {
			t.Errorf("Comparator.Compare() = %v, want %v", result, tt.expected)
		}
	}
}
