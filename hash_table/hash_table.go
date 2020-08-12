package hash_table

import (
	"bytes"
	"encoding/gob"
	"hash"
	"math/big"
)

type HashTable interface {
	Init(capacity uint32)
	Put(key interface{}, value interface{})
	Get(key interface{}) (interface{}, bool)
	Delete(key interface{})

	UpScale()
	Move(capacity uint32)
}

const (
	defaultCapacity   = 8
	defaultLoadFactor = 0.75
)

type hashTableBase struct {
	HashTable
	Capacity uint32
	Size     uint32
}

func (h *hashTableBase) Init(capacity uint32) {
	h.Capacity = capacity
	h.Size = 0
}

func (h *hashTableBase) CalculateLoad() float64 {
	if h.Capacity == 0 {
		return 1.0
	}
	return float64(h.Size) / float64(h.Capacity)
}

func (h *hashTableBase) HashFunc(key interface{}, hash hash.Hash) *big.Int {
	buf := bytes.NewBuffer(nil)
	enc := gob.NewEncoder(buf)
	_ = enc.Encode(key)
	hashBytes := hash.Sum(buf.Bytes())
	return new(big.Int).SetBytes(hashBytes)
}

func (h *hashTableBase) UpScale() {
	if h.CalculateLoad() >= defaultLoadFactor {
		if h.Capacity == 0 {
			h.HashTable.Init(defaultCapacity)
		} else {
			h.Move(h.Capacity << 1)
		} // else>>
	} // if>
}
