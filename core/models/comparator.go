package models

import "slices"

type Comparator struct {
}

func (e *Comparator) Compare(fact, value any) int {
	switch f := any(fact).(type) {
	case int:
		if f > any(value).(int) {
			return 1
		} else if f < any(value).(int) {
			return -1
		}
		return 0
	case string:
		if f > any(value).(string) {
			return 1
		} else if f < any(value).(string) {
			return -1
		}
		return 0
	case bool:
		if f && !any(value).(bool) {
			return 1
		} else if !f && any(value).(bool) {
			return -1
		}
		return 0
	default:
		panic("unsupported type for comparison")
	}
}

func (e *Comparator) Contains(fact, value any) bool {
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
