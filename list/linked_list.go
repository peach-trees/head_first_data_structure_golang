package list

import (
	"bytes"
	"fmt"
	"head_first_data_structure_golang/common"
)

type LinkedListNode struct {
	data interface{}
	next *LinkedListNode
}

type LinkedList struct {
	first *LinkedListNode
	last  *LinkedListNode
	size  int
}

func (l *LinkedList) checkIndex(index int) bool {
	return index >= 0 && index < l.size
}

func (l *LinkedList) Get(index int) (interface{}, bool) {

}

func (l *LinkedList) Set(index int, value interface{}) {
	if !l.checkIndex(index) {
		return
	}
	cur := l.first
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
	if l.size == 1 {
		l.Reset()
		return
	}

	var pre, cur *LinkedListNode
	cur = l.first
	for i := 0; i < index && cur != nil; i++ {
		pre = cur
		cur = cur.next
	} // for>
	if cur == nil {
		return
	}

	if cur == l.first {
		l.first = l.first.next
	}
	if cur == l.last {
		l.last = pre
	}
	if pre != nil {
		pre.next = cur.next
	}
	cur = nil
	l.size--
}

func (l *LinkedList) Append(values ...interface{}) {

}

func (l *LinkedList) Prepend(values ...interface{}) {

}

func (l *LinkedList) Insert(index int, values ...interface{}) {
	if !l.checkIndex(index) {
		return
	}
	l.size += len(values)

	var pre, cur *LinkedListNode
	cur = l.first
	for i := 0; i < index; i++ {
		pre = cur
		cur = cur.next
	} // for>

	if pre == nil {
		nextBackUp := l.first
		for i := len(values) - 1; i >= 0; i-- {
			newNode := &LinkedListNode{data: values[i], next: nextBackUp}
			nextBackUp = newNode
		} // for>>
		l.first = nextBackUp
	} else {
		nextBackUp := pre.next
		for _, v := range values {
			newNode := &LinkedListNode{data: v}
			pre.next = newNode
			pre = newNode
		} // for>>
		pre.next = nextBackUp
	} // else>
}

func (l *LinkedList) IndexOf(value interface{}) int {
	if l.size == 0 {
		return -1
	}
	var index int
	for cur := l.first; cur != nil; cur = cur.next {
		if cur.data == value {
			return index
		} // if>>
		index++
	} // for>
	return -1
}

func (l *LinkedList) Sort(comparator common.Comparator) {
	if l.size <= 1 {
		return
	}
	values := l.Values()
	common.Sort(values, comparator)
	l.Reset()
	l.Append(values...)
}

func (l *LinkedList) IfEmpty() bool {
	return l.size == 0
}

func (l *LinkedList) Size() int {
	return l.size
}

func (l *LinkedList) Reset() {
	l.size, l.first, l.last = 0, nil, nil
}

func (l *LinkedList) Values() []interface{} {
	ret := make([]interface{}, 0, l.size)
	for cur := l.first; cur != nil; cur = cur.next {
		ret = append(ret, cur.data)
	}
	return ret
}

func (l *LinkedList) String() string {
	retBytes := bytes.NewBufferString("[LinkedList]")
	for cur := l.first; cur != nil; cur = cur.next {
		retBytes.WriteString(fmt.Sprintf("->%+v", cur.data))
	} // for>
	return retBytes.String()
}

func NewLinkedList(values ...interface{}) *LinkedList {
	return nil
}
