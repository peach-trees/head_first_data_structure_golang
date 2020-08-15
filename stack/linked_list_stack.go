package stack

import (
	"bytes"
	"fmt"
	"head_first_data_structure_golang/list"
)

type LinkedListStack struct {
	list *list.LinkedList
}

func NewLinkedListStack() *LinkedListStack {
	return &LinkedListStack{list: &list.LinkedList{}}
}

func (l *LinkedListStack) Push(value interface{}) {
	l.list.Prepend(value)
}

func (l *LinkedListStack) Pop() (interface{}, bool) {
	value, ok := l.list.Get(0)
	l.list.Delete(0)
	return value, ok
}

func (l *LinkedListStack) Peek() (interface{}, bool) {
	return l.list.Get(0)
}

func (l *LinkedListStack) IfEmpty() bool {
	return l.list.IfEmpty()
}

func (l *LinkedListStack) Size() int {
	return l.list.Size()
}

func (l *LinkedListStack) Reset() {
	l.list.Reset()
}

func (l *LinkedListStack) String() string {
	retBytes := bytes.NewBufferString("[stack]")
	values := l.list.Values()
	for i := range values {
		retBytes.WriteString(fmt.Sprintf("->%v", values[i]))
	}
	return retBytes.String()
}
