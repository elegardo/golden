package core

import (
	"slices"

	"github.com/elegardo/golden/core/domain"
)

type FactComparator struct {
}

func (e *FactComparator) Compare(pair *domain.Pair) int {
	switch value1 := any(pair.Value1).(type) {
	case int:
		if value1 > pair.GetInt2() {
			return 1
		} else if value1 < pair.GetInt2() {
			return -1
		}
		return 0
	case string:
		if value1 > pair.GetString2() {
			return 1
		} else if value1 < pair.GetString2() {
			return -1
		}
		return 0
	case bool:
		if value1 && !pair.GetBool2() {
			return 1
		} else if !value1 && pair.GetBool2() {
			return -1
		}
		return 0
	default:
		panic("unsupported type for comparison")
	}
}

func (e *FactComparator) Contains(pair *domain.Pair) bool {
	switch value2 := any(pair.Value2).(type) {
	case interface{}:
		switch array := value2.(type) {
		case []string:
			return slices.Contains(array, pair.GetString1())
		case []int:
			return slices.Contains(array, pair.GetInt1())
		default:
			panic("unsupported type for contains")
		}
	default:
		panic("unsupported type for contains")
	}
}
