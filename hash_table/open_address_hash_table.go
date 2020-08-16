package hash_table

import (
	"crypto/sha1"
	"crypto/sha256"
	"head_first_data_structure_golang/common"
	"math/big"
)

type openAddressNode struct {
	Key    interface{}
	Value  interface{}
	exists bool
}

type OpenAddressHashTable struct {
	buckets  []*openAddressNode
	capacity uint32
	size     uint32
}

func NewOpenAddressHashTable(capacity uint32) *OpenAddressHashTable {
	h := new(OpenAddressHashTable)
	if capacity < defaultMinCapacity {
		h.capacity = defaultMinCapacity
		h.buckets = make([]*openAddressNode, defaultMinCapacity)
	} else {
		h.capacity = capacity
		h.buckets = make([]*openAddressNode, capacity)
	} // else>
	return h
}

// --- inner func ---
func (o *OpenAddressHashTable) hash(key interface{}, i uint32) uint32 {
	hashValue1, hashValue2 := common.HashFunc(key, sha1.New()), common.HashFunc(key, sha256.New())
	indexValue := big.NewInt(int64(i))
	capacityValue := big.NewInt(int64(o.capacity))
	hashValue2.Mul(hashValue2, indexValue).Add(hashValue2, hashValue1).Mod(hashValue2, capacityValue)
	return uint32(hashValue2.Uint64())
}

func (o *OpenAddressHashTable) UpScale() {
	if o.LoadFactor() >= defaultMinLoadFactor {
		oldBuckets := o.buckets
		o.capacity = o.capacity << 1
		o.buckets = make([]*openAddressNode, o.capacity)
		for i := range oldBuckets {
			if oldBuckets[i] != nil {
				o.Put(oldBuckets[i].Key, oldBuckets[i].Value)
			} // if>>>
		} // for>>
	} // if>
}

// --- interface func ---
func (o *OpenAddressHashTable) Put(key interface{}, value interface{}) {
	o.UpScale()
	for i := uint32(0); i < o.capacity; i++ {
		hashValue := o.hash(key, i)
		if o.buckets[hashValue] == nil {
			o.buckets[hashValue] = &openAddressNode{exists: false}
		} // if>>
		if o.buckets[hashValue] == nil || !o.buckets[hashValue].exists {
			// new
			o.buckets[hashValue].Key = key
			o.buckets[hashValue].Value = value
			o.buckets[hashValue].exists = true
			o.size++
			return
		} else if o.buckets[hashValue] != nil || o.buckets[hashValue].exists && o.buckets[hashValue].Key == key {
			// update
			o.buckets[hashValue].Value = value
			return
		} // else>>
	} // for>
}

func (o *OpenAddressHashTable) Get(key interface{}) (interface{}, bool) {
	if o.size == 0 {
		return nil, false
	}
	for i := uint32(0); i < o.capacity; i++ {
		hashValue := o.hash(key, i)
		if o.buckets[hashValue] != nil && o.buckets[hashValue].Key == key {
			return o.buckets[hashValue].Value, o.buckets[hashValue].exists
		} // if>>
	} // for>
	return nil, false
}

func (o *OpenAddressHashTable) Delete(key interface{}) {
	for i := uint32(0); i < o.capacity; i++ {
		hashValue := o.hash(key, i)
		if o.buckets[hashValue] != nil && o.buckets[hashValue].Key == key {
			o.buckets[hashValue] = &openAddressNode{exists: false}
			o.size--
			return
		} // if>>
	} // for>
}

func (o *OpenAddressHashTable) LoadFactor() float64 {
	if o.capacity == 0 {
		return 1.0
	}
	return float64(o.size) / float64(o.capacity)
}

func (o *OpenAddressHashTable) IfEmpty() bool {
	return o.size == 0
}

func (o *OpenAddressHashTable) Size() int {
	return int(o.size)
}

func (o *OpenAddressHashTable) Reset() {
	o.buckets = make([]*openAddressNode, o.capacity)
	o.size = 0
}
