package core

import (
	"slices"

	"github.com/elegardo/golden/core/interfaces"
)

type FactContainer[T interfaces.Value] struct {
}

// TODO: validar que sean todos los items del array del mismo tipo (Generics?)
func (e *FactContainer[T]) Contains(fact T, value []T) bool {
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
