package queue

import (
	"strings"
	"testing"
)

func TestNewLinkedQueue(t *testing.T) {
	aq := NewLinkedQueue()
	aq.Enqueue(3)
	aq.Enqueue(2)
	aq.Enqueue(1)
	t.Log(aq.String())

	for !aq.IfEmpty() {
		t.Log(strings.Repeat("-", 30))
		t.Log("head:")
		t.Log(aq.Head())
		t.Log("tail:")
		t.Log(aq.Tail())
		v, _ := aq.Dequeue()
		t.Log("CircularQueue:", v)
		t.Log(aq.String())
	}
	t.Log(strings.Repeat("-", 30))
	t.Log("head:")
	t.Log(aq.Head())
	t.Log("tail:")
	t.Log(aq.Tail())
	v, _ := aq.Dequeue()
	t.Log("CircularQueue:", v)
	t.Log(aq.String())

	t.Log(strings.Repeat("-", 30))
	aq.Enqueue(5)
	aq.Enqueue(6)
	aq.Enqueue(7)
	t.Log(aq.String())
}
