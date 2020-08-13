package list

import "head_first_data_structure_golang/common"

type List interface {
	Get(index int) (interface{}, bool)
	Set(index int, value interface{})
	Delete(index int)
	Append(values ...interface{})
	Prepend(values ...interface{})
	Insert(index int, values ...interface{})
	IndexOf(value interface{}) int
	Sort(comparator common.Comparator)
	Empty() bool
	Size() int
	Clear()
	String() string
}
