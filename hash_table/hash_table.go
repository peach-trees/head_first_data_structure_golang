package hash_table

type HashTable interface {
	Put(key interface{}, value interface{})
	Get(key interface{}) (interface{}, bool)
	Delete(key interface{})
	LoadFactor() float64
	Keys() []interface{}
	Values() []interface{}
	IfEmpty() bool
	Size() int
	Reset()
}

const (
	defaultMinCapacity   = 13
	defaultMinLoadFactor = 0.75
)
