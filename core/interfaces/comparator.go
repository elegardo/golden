package interfaces

import "github.com/elegardo/golden/core/domain"

type Comparator interface {
	Compare(pair *domain.Pair) int
	Contains(pair *domain.Pair) bool
}
