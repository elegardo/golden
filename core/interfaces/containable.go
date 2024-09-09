package interfaces

type Containable interface {
	Contains(Containable) bool
}
