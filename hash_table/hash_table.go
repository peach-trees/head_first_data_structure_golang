package hash_table

import "head_first_data_structure_golang/common"

type HashTable interface {
	common.Container
	Put(key interface{}, value interface{})
	Get(key interface{}) (interface{}, bool)
	Delete(key interface{})
	Keys() []interface{}
}
