package hash_table

import (
	"bytes"
	"encoding/gob"
	"hash"
	"math/big"
)

const (
	DefaultCapacity   = 8
	DefaultLoadFactor = 0.75
)

type ScaleAbleHashTable interface {
	HashTable
	UpScale()
	DownScale()
	Move(uint32)
}

type HashTableBase struct {
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
