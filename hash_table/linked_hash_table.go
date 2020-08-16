package hash_table

import (
	"container/list"
	"crypto/sha1"
	"head_first_data_structure_golang/common"
	"math/big"
)

type linkedNode struct {
	Key   interface{}
	Value interface{}
}

type LinkedHashTable struct {
	buckets  []*list.List
	capacity uint32
	size     uint32
}

func NewLinkedHashTable(capacity uint32) *LinkedHashTable {
	h := &LinkedHashTable{}
	if capacity < defaultMinCapacity {
		h.capacity = defaultMinCapacity
		h.buckets = make([]*list.List, defaultMinCapacity)
	} else {
		h.capacity = capacity
		h.buckets = make([]*list.List, capacity)
	} // else>
	return h
}

// --- inner func ---
func (l *LinkedHashTable) hash(key interface{}) uint32 {
	hashValue := common.HashFunc(key, sha1.New())
	capacityValue := big.NewInt(int64(l.capacity))
	hashValue.Mod(hashValue, capacityValue)
	return uint32(hashValue.Uint64())
}

func (l *LinkedHashTable) findInList(key interface{}, list *list.List) (*list.Element, bool) {
	for element := list.Front(); element != nil; element = element.Next() {
		if element.Value.(linkedNode).Key == key {
			return element, true
		} // if>>
	} // for>
	return nil, false
}

func (l *LinkedHashTable) Put(key interface{}, value interface{}) {
	hashKey := l.hash(key)
	if l.buckets[hashKey] == nil {
		l.buckets[hashKey] = list.New()
	} // if>

	element := linkedNode{Key: key, Value: value}
	listElement, exists := l.findInList(key, l.buckets[hashKey])
	if exists {
		listElement.Value = element
	} else {
		l.buckets[hashKey].PushFront(element)
		l.size++
	} // else>
}

func (l *LinkedHashTable) Get(key interface{}) (interface{}, bool) {
	if l.size == 0 {
		return nil, false
	}
	hashKey := l.hash(key)
	if l.buckets[hashKey] == nil {
		return nil, false
	}
	listElement, exists := l.findInList(hashKey, l.buckets[hashKey])
	if exists {
		return listElement.Value.(linkedNode).Value, true
	}
	return nil, false
}

func (l *LinkedHashTable) Delete(key interface{}) {
	hashKey := l.hash(key)
	if l.buckets[hashKey] == nil {
		return
	}
	listElement, exists := l.findInList(key, l.buckets[hashKey])
	if exists {
		l.buckets[hashKey].Remove(listElement)
		l.size--
	}
	if l.buckets[hashKey].Len() == 0 {
		l.buckets[hashKey] = nil
	}
}

func (l *LinkedHashTable) LoadFactor() float64 {
	if l.capacity == 0 {
		return 1.0
	}
	return float64(l.size) / float64(l.capacity)
}

func (l *LinkedHashTable) Keys() []interface{} {
	keys := make([]interface{}, 0, l.size)
	for _, v := range l.buckets {
		if v != nil {
			for element := v.Front(); element != nil; element = element.Next() {
				keys = append(keys, element.Value.(linkedNode).Key)
			} // for>>
		} // if>>
	} // for>
	return keys
}

func (l *LinkedHashTable) Values() []interface{} {
	values := make([]interface{}, 0, l.size)
	for _, v := range l.buckets {
		if v != nil {
			for element := v.Front(); element != nil; element = element.Next() {
				values = append(values, element.Value.(linkedNode).Value)
			} // for>>
		} // if>>
	} // for>
	return values
}

func (l *LinkedHashTable) IfEmpty() bool {
	return l.size == 0
}

func (l *LinkedHashTable) Size() int {
	return int(l.size)
}

func (l *LinkedHashTable) Reset() {
	l.buckets = make([]*list.List, l.capacity)
	l.size = 0
}
