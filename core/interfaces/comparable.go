// Paquete de interfaces usadas para construir los distintos tipos de Engine.
package interfaces

// Permite comparar 2 valores del mismo tipo. Por ejemplo:
//
//	apple  := Fruit{Value: "apple"}
//	banana := Fruit{Value: "banana"}
//	apple.Compare(banana) // false
//
// El tipo debe estar definido previamente y puede ser cualquiera que use el generico [T comparable].
//
//	...
//	type Age[T comparable] struct {
//		Value T
//	}
//	...
//	number1  := Age{Value: 10}
//	number2  := Age{Value: 10}
//	number1.Compare(number2) // true
type Comparable interface {
	Compare(Comparable) int
}
