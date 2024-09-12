package core

import (
	"slices"

	"github.com/elegardo/golden/core/interfaces"
)

type Container[T comparable] struct {
	Value T
}

func (value Container[T]) Contains(fact interfaces.Containable) bool {
	y := fact.(Container[T])
	switch x := any(value.Value).(type) {
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
