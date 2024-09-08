package engine

type Event struct {
	order int
	text  string
}

func (e Event) Value() any {
	return e.text
}

func (e Event) Order() int {
	return e.order
}
