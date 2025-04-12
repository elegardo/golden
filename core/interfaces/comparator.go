package interfaces

type Comparator interface {
	Compare(fact, value any) int
	Contains(fact, value any) bool
}
