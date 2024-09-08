package engine

import (
	"slices"
	"strings"
)

type Container[T comparable] struct {
	Value T
}

func (value Container[T]) Contains(fact Containable) bool {
	y := fact.(Container[T])
	switch x := any(value.Value).(type) {
	case string:
		return strings.Contains(x, any(y.Value).(string))
	case interface{}:
		switch array := x.(type) {
		case []string:
			return slices.Contains(array, any(y.Value).(string))
		case []int:
			return slices.Contains(array, any(y.Value).(int))
		default:
			panic("unsupported type for contains")
		}
	default:
		panic("unsupported type for contains")
	}
}
