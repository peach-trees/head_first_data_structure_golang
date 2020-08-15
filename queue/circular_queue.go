package queue

import (
	"bytes"
	"errors"
	"fmt"
)

const defaultMinCapacity = 8

type CircularQueue struct {
	head     int
	tail     int
	capacity int
	data     []interface{}
}

func NewCircularQueue(capacity int) *CircularQueue {
	if capacity <= 2 {
		capacity = defaultMinCapacity
	}
	return &CircularQueue{
		head:     0,
		tail:     0,
		capacity: capacity,
		data:     make([]interface{}, 0, capacity),
	}
}

func (d *CircularQueue) IfEmpty() bool {
	return d.head == d.tail
}

func (d *CircularQueue) IfFull() bool {
	return ((d.tail + 1) % d.capacity) == d.head
}

func (d *CircularQueue) Size() int {
	return (d.tail - d.head + d.capacity) % d.capacity
}

func (d *CircularQueue) Reset() {
	d.head, d.tail = 0, 0
	d.data = make([]interface{}, d.capacity, d.capacity)
}

func (d *CircularQueue) String() string {
	retBytes := bytes.NewBufferString("[Queue-head]")
	for i := d.head; i != d.tail; i = (i + 1) % d.capacity {
		retBytes.WriteString(fmt.Sprintf("->%+v", d.data[i]))
	}
	retBytes.WriteString("->[Queue-tail]")
	return retBytes.String()
}

func (d *CircularQueue) Enqueue(value interface{}) error {
	if d.IfFull() {
		return errors.New("CircularQueue is full")
	}
	d.data[d.tail] = value
	d.tail = (d.tail + 1) % d.capacity
	return nil
}

func (d *CircularQueue) Dequeue() (interface{}, bool) {
	if d.IfEmpty() {
		return nil, false
	}
	ret := d.data[d.head]
	d.head = (d.head + 1) % d.capacity
	return ret, true
}

func (d *CircularQueue) Head() (interface{}, bool) {
	if d.IfEmpty() {
		return nil, false
	}
	return d.data[d.head], true
}

func (d *CircularQueue) Tail() (interface{}, bool) {
	if d.IfEmpty() {
		return nil, false
	}
	return d.data[d.tail], true
}
