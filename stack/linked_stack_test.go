package stack

import "testing"

func TestNewLinkedStack(t *testing.T) {
	ls := NewLinkedStack()
	ls.Push(1)
	ls.Push(2)
	ls.Push(3)
	t.Log("size:", ls.size, ", string:", ls.String())

	t.Log("peek:")
	t.Log(ls.Peek())
	v, _ := ls.Pop()
	t.Log("pop:", v)
	t.Log("size:", ls.size, ", string:", ls.String())

	t.Log("peek:")
	t.Log(ls.Peek())
	v, _ = ls.Pop()
	t.Log("pop:", v)
	t.Log("size:", ls.size, ", string:", ls.String())

	t.Log("peek:")
	t.Log(ls.Peek())
	v, _ = ls.Pop()
	t.Log("pop:", v)
	t.Log("size:", ls.size, ", string:", ls.String())

	t.Log("peek:")
	t.Log(ls.Peek())
	v, _ = ls.Pop()
	t.Log("pop:", v)
	t.Log("size:", ls.size, ", string:", ls.String())
}
