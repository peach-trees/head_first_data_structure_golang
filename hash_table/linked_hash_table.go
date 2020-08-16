package hash_table

import "container/list"

type node struct {
	Key   interface{}
	Value interface{}
}

type LinkedHashTable struct {
	buckets  []*list.List
	capacity uint32
}

func NewLinkedHashTable(capacity uint32) *LinkedHashTable {
	h := &LinkedHashTable{capacity: capacity}
	if capacity == 0 {
		h.buckets = nil
	} else {
		h.buckets = make([]*list.List, capacity)
	} // else>
	return h
}

func (l *LinkedHashTable) Put(key interface{}, value interface{}) {

}

func (l *LinkedHashTable) Get(key interface{}) (interface{}, bool) {

}

func (l *LinkedHashTable) Delete(key interface{}) {

}

func (l *LinkedHashTable) Keys() []interface{} {

}

func (l *LinkedHashTable) Values() []interface{} {

}

func (l *LinkedHashTable) IfEmpty() bool {

}

func (l *LinkedHashTable) Size() int {

}

func (l *LinkedHashTable) Reset() {
	if l.capacity == 0 {
		l.buckets = nil
	} else {
		l.buckets = make([]*list.List, l.capacity)
	} // else>
}
