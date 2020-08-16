package list

import (
	"bytes"
	"fmt"
)

type LinkedListNode struct {
	data interface{}
	next *LinkedListNode
}

type LinkedList struct {
	head *LinkedListNode
	size int
}

func NewLinkedList(values ...interface{}) *LinkedList {
	list := &LinkedList{
		head: nil,
		size: 0,
	}
	if len(values) > 0 {
		list.Append(values...)
	}
	return list
}

func (l *LinkedList) checkIndex(index int) bool {
	return index >= 0 && index < l.size
}

func (l *LinkedList) Append(values ...interface{}) {
	tail := l.head
	for tail != nil && tail.next != nil {
		tail = tail.next
	} // for>

	for i := range values {
		newNode := &LinkedListNode{data: values[i]}
		if tail == nil {
			l.head = newNode
			tail = newNode
		} else {
			tail.next = newNode
			tail = newNode
		} // else>>
		l.size++
	} // for>
}

func (l *LinkedList) Prepend(values ...interface{}) {
	for i := len(values) - 1; i >= 0; i-- {
		newNode := &LinkedListNode{data: values[i], next: l.head}
		l.head = newNode
		l.size++
	} // for>
}

func (l *LinkedList) Get(index int) (interface{}, bool) {
	if !l.checkIndex(index) {
		return nil, false
	}
	cur := l.head
	for i := 0; i < index && cur != nil; i++ {
		cur = cur.next
	}
	if cur != nil {
		return cur.data, true
	}
	return nil, false
}

func (l *LinkedList) Set(index int, value interface{}) {
	if !l.checkIndex(index) {
		return
	}
	cur := l.head
	for i := 0; i < index && cur != nil; i++ {
		cur = cur.next
	}
	if cur != nil {
		cur.data = value
	}
}

func (l *LinkedList) Delete(index int) {
	if !l.checkIndex(index) {
		return
	}

	var pre, cur *LinkedListNode
	cur = l.head
	for i := 0; i < index && cur != nil; i++ {
		pre = cur
		cur = cur.next
	} // for>
	if cur == nil {
		return
	}

	if pre == nil {
		l.head = l.head.next
	} else {
		pre.next = cur.next
	} // else>
	l.size--
}

func (l *LinkedList) Insert(index int, value interface{}) {
	if !l.checkIndex(index) {
		return
	}

	cur := l.head
	for i := 0; i < index-1 && cur != nil; i++ {
		cur = cur.next
	} // for>

	if cur == nil {
		newNode := &LinkedListNode{data: value, next: l.head}
		l.head = newNode
	} else {
		newNode := &LinkedListNode{data: value, next: cur.next}
		cur.next = newNode
	} // else>

	l.size += 1
}

func (l *LinkedList) IndexOf(value interface{}) int {
	if l.size == 0 {
		return -1
	}
	var index int
	for cur := l.head; cur != nil; cur = cur.next {
		if cur.data == value {
			return index
		} // if>>
		index++
	} // for>
	return -1
}

func (l *LinkedList) IfEmpty() bool {
	return l.size == 0
}

func (l *LinkedList) Size() int {
	return l.size
}

func (l *LinkedList) Reset() {
	l.size, l.head = 0, nil
}

func (l *LinkedList) String() string {
	retBytes := bytes.NewBufferString("[LinkedList]")
	for cur := l.head; cur != nil; cur = cur.next {
		retBytes.WriteString(fmt.Sprintf("->%+v", cur.data))
	} // for>
	return retBytes.String()
}
