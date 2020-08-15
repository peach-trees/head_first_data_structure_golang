package stack

import "testing"

func TestNewArrayStack(t *testing.T) {
	s := NewArrayStack()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	t.Log(s.String())

	v, _ := s.Pop()
	t.Log("pop:", v)
	t.Log(s.String())

	v, _ = s.Pop()
	t.Log("pop:", v)
	t.Log(s.String())

	v, _ = s.Pop()
	t.Log("pop:", v)
	t.Log(s.String())

	v, _ = s.Pop()
	t.Log("pop:", v)
	t.Log(s.String())
}
