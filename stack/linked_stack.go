package stack

import (
	"bytes"
	"fmt"
)

type node struct {
	data interface{}
	next *node
}

type LinkedStack struct {
	head *node
	size int
}

func NewLinkedStack() *LinkedStack {
	return &LinkedStack{}
}

func (l *LinkedStack) Push(value interface{}) {
	newNode := &node{data: value, next: nil}
	if l.head == nil {
		l.head = newNode
	} else {
		newNode.next = l.head
		l.head = newNode
	} // else>
	l.size++
}

func (l *LinkedStack) Pop() (interface{}, bool) {
	if l.head == nil {
		return nil, false
	} // if>
	ret := l.head.data
	l.head = l.head.next
	l.size--
	return ret, true
}

func (l *LinkedStack) Peek() (interface{}, bool) {
	if l.head == nil {
		return nil, false
	} // if>
	return l.head.data, true
}

func (l *LinkedStack) IfEmpty() bool {
	return l.head == nil
}

func (l *LinkedStack) Size() int {
	return l.size
}

func (l *LinkedStack) Reset() {
	l.head, l.size = nil, 0
}

func (l *LinkedStack) String() string {
	retBytes := bytes.NewBufferString("[Stack-top]")
	for cur := l.head; cur != nil; cur = cur.next {
		retBytes.WriteString(fmt.Sprintf("->%+v", cur.data))
	} // for>
	retBytes.WriteString("->[Stack-bottom]")
	return retBytes.String()
}
