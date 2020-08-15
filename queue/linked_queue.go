package queue

import (
	"bytes"
	"fmt"
)

type node struct {
	data interface{}
	next *node
}

type LinkedQueue struct {
	head *node
	tail *node
	size int
}

func NewLinkedQueue() *LinkedQueue {
	return &LinkedQueue{}
}

func (l *LinkedQueue) IfEmpty() bool {
	return l.size == 0
}

func (l *LinkedQueue) Size() int {
	return l.size
}

func (l *LinkedQueue) Reset() {
	l.head, l.tail, l.size = nil, nil, 0
}

func (l *LinkedQueue) String() string {
	retBytes := bytes.NewBufferString("[Queue-head]")
	for cur := l.head; cur != nil; cur = cur.next {
		retBytes.WriteString(fmt.Sprintf("->%+v", cur.data))
	}
	retBytes.WriteString("->[Queue-tail]")
	return retBytes.String()
}

func (l *LinkedQueue) Enqueue(value interface{}) {
	l.size++
	newNode := &node{data: value}
	if l.head == nil && l.tail == nil {
		l.head = newNode
		l.tail = newNode
	} else {
		l.tail.next = newNode
		l.tail = newNode
	}
}

func (l *LinkedQueue) Dequeue() (interface{}, bool) {
	if l.head == nil && l.tail == nil {
		return nil, false
	}
	l.size--
	ret := l.head.data

	l.head = l.head.next
	if l.head == nil {
		l.tail = nil
	}

	return ret, true
}

func (l *LinkedQueue) Head() (interface{}, bool) {
	if l.head == nil && l.tail == nil {
		return nil, false
	}
	return l.head.data, true
}

func (l *LinkedQueue) Tail() (interface{}, bool) {
	if l.head == nil && l.tail == nil {
		return nil, false
	}
	return l.tail.data, true
}
