package core

import "github.com/elegardo/golden/core/interfaces"

// Permite comparar 2 valores del tipo Comparator. El generico [T comparable] puede ser de cualquier tipo
type Comparator[T comparable] struct {
	Value T
}

// Por ejemplo:
//
//	fact  := Comparator[any]{Value: "apple"}
//	value := Comparator[any]{Value: "banana"}
//	fact.Compare(value) // false
//
// Los valores pueden ser de cualquier tipo, pero deben ser del mismo tipo:
//
//	fact  := Comparator[any]{Value: 10}
//	value := Comparator[any]{Value: 10}
//	fact.Compare(value) // true
func (fact Comparator[T]) Compare(value interfaces.Comparable) int {
	v := value.(Comparator[T]).Value
	switch f := any(fact.Value).(type) {
	case int:
		if f > any(v).(int) {
			return 1
		} else if f < any(v).(int) {
			return -1
		}
		return 0
	case string:
		if f > any(v).(string) {
			return 1
		} else if f < any(v).(string) {
			return -1
		}
		return 0
	case bool:
		if f && !any(v).(bool) {
			return 1
		} else if !f && any(v).(bool) {
			return -1
		}
		return 0
	default:
		panic("unsupported type for comparison")
	}
}
