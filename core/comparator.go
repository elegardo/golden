package core

import "github.com/elegardo/golden/core/interfaces"

// Estructura que permite comparar 2 valores, un "fact" y un "value" de una regla
type Comparator[T comparable] struct {
	Value T
}

// La función permite comparar 2 valores. Por ejemplo:
//
//	apple  := Comparator[any]{Value: "apple"}
//	banana := Comparator[any]{Value: "banana"}
//	apple.Compare(banana) // false
//
// El valor de "result" es false.
// Los valores deben ser del mismo tipo:
//
//	number1  := Comparator[any]{Value: 10}
//	number2  := Comparator[any]{Value: 10}
//	apple.Compare(banana) // true
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
