package interfaces

type Comparable interface {
	Compare(fact, value any) int
	Contains(fact, value any) bool
}
