package hash_table

import (
	"crypto/sha1"
	"crypto/sha256"
	"math/big"
)

type openAddressElement struct {
	Key    interface{}
	Value  interface{}
	exists bool
}

type OpenAddressHashTable struct {
	hashTableBase
	buckets []*openAddressElement
}

func NewOpenAddressHashTable() *OpenAddressHashTable {
	h := new(OpenAddressHashTable)
	h.hashTableBase.HashTable = h
	return h
}

// --- inner func ---
func (o *OpenAddressHashTable) hash(key interface{}, i uint32) uint32 {
	hashValue1, hashValue2 := o.HashFunc(key, sha1.New()), o.HashFunc(key, sha256.New())
	ib := big.NewInt(int64(i))
	mb := big.NewInt(int64(o.Capacity))
	hashValue2.Mul(hashValue2, ib).Add(hashValue2, hashValue1).Mod(hashValue2, mb)
	return uint32(hashValue2.Uint64())
}

func (o *OpenAddressHashTable) ifExists(key uint32) bool {
	if o.buckets[key] == nil {
		return false
	}
	return o.buckets[key].exists
}

// --- interface func ---
func (o *OpenAddressHashTable) Init(capacity uint32) {
	o.hashTableBase.Init(capacity)
	o.buckets = make([]*openAddressElement, 0, o.Capacity)
}

func (o *OpenAddressHashTable) Move(capacity uint32) {
	oldBuckets := o.buckets
	o.Init(capacity)
	for i := range oldBuckets {
		if oldBuckets[i] != nil {
			o.Put(oldBuckets[i].Key, oldBuckets[i].Value)
		} // if>>
	} // for>
}

func (o *OpenAddressHashTable) Put(key interface{}, value interface{}) {
	o.UpScale()
	for i := 0; i < int(o.Capacity); i++ {
		hashValue := o.hash(key, uint32(i))
		if o.buckets[hashValue] == nil {
			o.buckets[hashValue] = &openAddressElement{exists: false}
		} // if>
		exists := o.ifExists(hashValue)
		if !exists {
			o.buckets[hashValue].Key = key
			o.buckets[hashValue].Value = value
			o.buckets[hashValue].exists = true
			o.Size++
			return
		} else if exists && o.buckets[hashValue].Key == key {
			o.buckets[hashValue].Value = value
			return
		}
	} // for>
}

func (o *OpenAddressHashTable) Get(key interface{}) (interface{}, bool) {

}

func (o *OpenAddressHashTable) Delete(key interface{}) {

}
