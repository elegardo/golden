package core

type Comparator[T comparable] struct {
	Value T
}

func (fact Comparator[T]) Compare(value Comparable) int {
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
