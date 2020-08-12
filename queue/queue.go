package queue

type Queue interface {
	Size() int
	IsEmpty() bool
	PushBack(interface{})
	PeekBack() interface{}
	PopFront() interface{}
	PeekFront() interface{}
}

const (
	defaultCapacity = 8
)
