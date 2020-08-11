package linked_list_hash_table

import (
	"container/list"
	"head_first_data_structure_golang/hash_table"
)

type LinkedListHashTable struct {
	hash_table.HashTableBase
	buckets []*list.List
}

func (l *LinkedListHashTable) Init(capacity uint32) {
	l.HashTableBase.Init(capacity)
	if capacity == 0 {
		l.buckets = nil
	} else {
		l.buckets = make([]*list.List, l.Capacity)
	} // else>
}

func (l *LinkedListHashTable) Move(capacity uint32) {
	oldBuckets := l.buckets
	l.Init(capacity)
	for _, tempList := range oldBuckets {
		if tempList != nil {
			for cur := tempList.Front(); cur != nil; cur = cur.Next() {
				l.Put(cur.Value.(hash_table.HashTableElement).Key, cur.Value.(hash_table.HashTableElement).Value)
			} // for>>>
		} // if>>
	} // for>
}

func (l *LinkedListHashTable) Put(key interface{}, value interface{}) {

}
