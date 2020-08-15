package stack

type Stack interface {
	Push(interface{})
	Pop() (interface{}, bool)
	Peek() (interface{}, bool)
	IfEmpty() bool
	Size() int
	Reset()
	String() string
}
