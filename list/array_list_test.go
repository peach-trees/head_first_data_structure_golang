package list

import (
	"head_first_data_structure_golang/common"
	"testing"
)

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

func TestArrayList_Sort(t *testing.T) {
	tl := NewArrayList(1, 2, 3)
	tl.Append(11, 10, 9)
	tl.Prepend(5, 6, 7)
	t.Log(tl.String())
	tl.Sort(common.IntComparator)
	t.Log(tl.String())
}

func TestArrayList_Insert(t *testing.T) {
	tl := NewArrayList(1, 2, 3)
	t.Log(tl.String())
	tl.Insert(1, 5, 6)
	t.Log(tl.String())
	tl.Insert(0, 7, 8)
	t.Log(tl.String())
}
