package engine

type ExampleEvent struct {
	order int
	text  string
}

func (e ExampleEvent) Value() any {
	return e.text
}

func (e ExampleEvent) Order() int {
	return e.order
}
