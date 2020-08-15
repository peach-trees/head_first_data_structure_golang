package stack

import "testing"

func TestNewLinkedListStack(t *testing.T) {
	s := NewLinkedListStack()
	s.Push(1)
	s.Push(2)
	t.Log(s.String())
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

}
