package hash_table

import (
	"bytes"
	"encoding/gob"
	"hash"
	"head_first_data_structure_golang/common"
	"math/big"
)

type HashTable interface {
	common.Container
	Init(capacity uint32)
	Put(key interface{}, value interface{})
	Get(key interface{}) (interface{}, bool)
	Delete(key interface{})
	Keys() []interface{}
}

const (
	DefaultCapacity   = 8
	DefaultLoadFactor = 0.75
)

type ScaleAbleHashTable interface {
	UpScale()
	DownScale()
	Move(uint32)
}

type HashTableBase struct {
	HashTable
	ScaleAbleHashTable
	Capacity uint32
	Size     uint32
}

func (h *HashTableBase) Init(capacity uint32) {
	h.Capacity = capacity
	h.Size = 0
}

func (h *HashTableBase) CalculateLoad() float64 {
	if h.Capacity == 0 {
		return 1.0
	}
	return float64(h.Size) / float64(h.Capacity)
}

func (h *HashTableBase) HashFunc(key interface{}, hash hash.Hash) *big.Int {
	buf := bytes.NewBuffer(nil)
	enc := gob.NewEncoder(buf)
	_ = enc.Encode(key)
	hashBytes := hash.Sum(buf.Bytes())
	return new(big.Int).SetBytes(hashBytes)
}

func (h *HashTableBase) UpScale() {
	if h.CalculateLoad() >= DefaultLoadFactor {
		if h.Capacity == 0 {
			h.HashTable.Init(DefaultCapacity)
		} else {
			h.Move(h.Capacity << 1)
		} // else>>
	} // if>
}
