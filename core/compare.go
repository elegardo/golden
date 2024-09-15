package core

func Compare(fact, value any) int {
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
