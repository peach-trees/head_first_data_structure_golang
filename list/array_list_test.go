package list

import "testing"

func TestArrayList_Append(t *testing.T) {
	tl := NewArrayList(1, 2, 3)
	t.Log(tl.String())

	tl.Append('a', 'b', 'c')
	t.Log(tl.String())

	tl.Append('x', 'y', 'z')
	t.Log(tl.String())
}

func TestArrayList_Prepend(t *testing.T) {
	tl := NewArrayList(1, 2, 3)
	t.Log(tl.String())

	tl.Prepend('a', 'b', 'c')
	t.Log(tl.String())

	tl.Prepend('x', 'y', 'z')
	t.Log(tl.String())
}
