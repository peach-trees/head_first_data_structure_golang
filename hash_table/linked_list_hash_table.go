package hash_table

import (
	"container/list"
	"crypto/sha1"
	"math/big"
)

type LinkedListHashTable struct {
	HashTableBase
	buckets []*list.List
}

// --- inner func ---
func (l *LinkedListHashTable) hash(key interface{}) uint32 {
	hashValue := l.HashFunc(key, sha1.New())
	mb := big.NewInt(int64(l.Capacity))
	hashValue.Mod(hashValue, mb)
	return uint32(hashValue.Uint64())
}

func (l *LinkedListHashTable) findInList(key interface{}, list *list.List) (*list.Element, bool) {
	for element := list.Front(); element != nil; element = element.Next() {
		if element.Value.(HashTableElement).Key == key {
			return element, true
		} // if>>
	} // for>
	return nil, false
}

// --- interface func ---
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
				l.Put(cur.Value.(HashTableElement).Key, cur.Value.(HashTableElement).Value)
			} // for>>>
		} // if>>
	} // for>
}

func (l *LinkedListHashTable) Put(key interface{}, value interface{}) {
	l.UpScale()
	hashKey := l.hash(key)
	if l.buckets[hashKey] == nil {
		l.buckets[hashKey] = list.New()
	} // if>
	element := HashTableElement{
		Key:   key,
		Value: value,
	}
	listElement, exists := l.findInList(key, l.buckets[hashKey])
	if exists {
		listElement.Value = element
	} else {
		l.buckets[hashKey].PushFront(element)
		l.Size++
	}
}

func (l *LinkedListHashTable) Get(key interface{}) (interface{}, bool) {

}

func (l *LinkedListHashTable) Delete(key interface{}) {

}

func (l *LinkedListHashTable) Keys() []interface{} {

}
