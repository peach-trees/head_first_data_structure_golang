package queue

import (
	"bytes"
	"fmt"
)

type ArrayQueue struct {
	data []interface{}
}

func NewArrayQueue() *ArrayQueue {
	return &ArrayQueue{data: make([]interface{}, 0)}
}

func (a *ArrayQueue) IfEmpty() bool {
	return len(a.data) == 0
}

func (a *ArrayQueue) Size() int {
	return len(a.data)
}

func (a *ArrayQueue) Reset() {
	a.data = make([]interface{}, 0)
}

func (a *ArrayQueue) String() string {
	retBytes := bytes.NewBufferString("[Queue-head]")
	for i := range a.data {
		retBytes.WriteString(fmt.Sprintf("->%+v", a.data[i]))
	}
	retBytes.WriteString("->[Queue-tail]")
	return retBytes.String()
}

func (a *ArrayQueue) Enqueue(value interface{}) {
	a.data = append(a.data, value)
}

func (a *ArrayQueue) Dequeue() (interface{}, bool) {
	if len(a.data) == 0 {
		return nil, false
	}
	ret := a.data[0]
	a.data = a.data[1:]
	return ret, true
}

func (a *ArrayQueue) Head() (interface{}, bool) {
	if len(a.data) == 0 {
		return nil, false
	}
	return a.data[0], true
}

func (a *ArrayQueue) Tail() (interface{}, bool) {
	if len(a.data) == 0 {
		return nil, false
	}
	return a.data[len(a.data)-1], true
}
