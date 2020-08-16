package hash_table

import (
	"bytes"
	"container/list"
	"crypto/sha1"
	"encoding/gob"
	"math/big"
)

const defaultMinCapacity = 8

type node struct {
	Key   interface{}
	Value interface{}
}

type LinkedHashTable struct {
	buckets  []*list.List
	capacity uint32
	size     uint32
}

func NewLinkedHashTable(capacity uint32) *LinkedHashTable {
	h := &LinkedHashTable{capacity: capacity, size: 0}
	if capacity == 0 {
		h.buckets = nil
	} else {
		h.buckets = make([]*list.List, capacity)
	} // else>
	return h
}

// --- inner func ---
func (l *LinkedHashTable) hash(key interface{}) uint32 {
	buf := bytes.NewBuffer(nil)
	enc := gob.NewEncoder(buf)
	_ = enc.Encode(key)
	hashMethod := sha1.New()
	hashBytes := hashMethod.Sum(buf.Bytes())
	hashValue := new(big.Int).SetBytes(hashBytes)

	capacityValue := big.NewInt(int64(l.capacity))
	hashValue.Mod(hashValue, capacityValue)
	return uint32(hashValue.Uint64())
}

func (l *LinkedHashTable) findInList(key interface{}, list *list.List) (*list.Element, bool) {
	for element := list.Front(); element != nil; element = element.Next() {
		if element.Value.(node).Key == key {
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

	element := node{Key: key, Value: value}
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
		return listElement.Value.(node).Value, true
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
	}
	if l.buckets[hashKey].Len() == 0 {
		l.buckets[hashKey] = nil
		l.size--
	}
}

func (l *LinkedHashTable) Load() float64 {
	if l.capacity == 0 {
		return 1.0
	}
	return float64(l.size) / float64(l.capacity)
}

func (l *LinkedHashTable) Keys() []interface{} {
	keys := make([]interface{}, 0, s.Size())
	for k := range s.data {
		keys = append(keys, k)
	}
	return keys
}

func (l *LinkedHashTable) Values() []interface{} {
	values := make([]interface{}, 0, s.Size())
	for _, v := range s.data {
		values = append(values, v)
	}
	return values
}

func (l *LinkedHashTable) IfEmpty() bool {
	return l.size == 0
}

func (l *LinkedHashTable) Size() int {
	return int(l.size)
}

func (l *LinkedHashTable) Reset() {
	if l.capacity == 0 {
		l.buckets = nil
	} else {
		l.buckets = make([]*list.List, l.capacity)
	} // else>
}
