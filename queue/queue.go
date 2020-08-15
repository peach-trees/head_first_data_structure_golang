package queue

type Queue interface {
	IfEmpty() bool
	Size() int
	Reset()
	String() string
	Enqueue(interface{})
	Dequeue() (interface{}, bool)
	Head() (interface{}, bool)
	Tail() (interface{}, bool)
}

const (
	defaultCapacity = 8
)
