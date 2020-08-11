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
