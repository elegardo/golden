package core

import (
	"testing"
)

func TestComparator_Compare(t *testing.T) {
	tests := []struct {
		name     string
		a        Comparator[any]
		b        Comparator[any]
		expected int
	}{
		{
			name:     "Compare equal integers",
			a:        Comparator[any]{Value: 10},
			b:        Comparator[any]{Value: 10},
			expected: 0,
		},
		{
			name:     "Compare integers a > b",
			a:        Comparator[any]{Value: 10},
			b:        Comparator[any]{Value: 5},
			expected: 1,
		},
		{
			name:     "Compare integers a < b",
			a:        Comparator[any]{Value: 5},
			b:        Comparator[any]{Value: 10},
			expected: -1,
		},
		{
			name:     "Compare equal strings",
			a:        Comparator[any]{Value: "apple"},
			b:        Comparator[any]{Value: "apple"},
			expected: 0,
		},
		{
			name:     "Compare strings a > b",
			a:        Comparator[any]{Value: "banana"},
			b:        Comparator[any]{Value: "apple"},
			expected: 1,
		},
		{
			name:     "Compare strings a < b",
			a:        Comparator[any]{Value: "apple"},
			b:        Comparator[any]{Value: "banana"},
			expected: -1,
		},
	}

	for _, tt := range tests {
		result := tt.a.Compare(tt.b)
		if result != tt.expected {
			t.Errorf("Comparator.Compare() = %v, want %v", result, tt.expected)
		}
	}
}
