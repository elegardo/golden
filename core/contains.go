package core

import (
	"slices"
)

func Contains(fact, value any) bool {
	switch x := any(value).(type) {
	case interface{}:
		switch array := x.(type) {
		case []string:
			return slices.Contains(array, any(fact).(string))
		case []int:
			return slices.Contains(array, any(fact).(int))
		default:
			panic("unsupported type for contains")
		}
	default:
		panic("unsupported type for contains")
	}
}
