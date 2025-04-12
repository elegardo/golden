package interfaces

type Container interface {
	Contains(fact any, value []any) bool
}
