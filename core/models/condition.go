package models

type Condition struct {
	Gate         Gate
	Conditionals []Conditional
}

// Priorizar ejecución de Gates en el siguiente orden: ALL, NONE, ANY.
// Lo anterior permite descartar mas rapido un false,
type ConditionSorter []Condition

func (a ConditionSorter) Len() int      { return len(a) }
func (a ConditionSorter) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ConditionSorter) Less(i, j int) bool {
	return a[i].Gate == ANY && a[j].Gate != ANY
}
