package standard_hash_table

type HashTable struct {
	data map[interface{}]interface{}
}

func NewHashTable(capacity uint) *HashTable {
	ret := &HashTable{}
	ret.Init(capacity)
	return ret
}

func (h *HashTable) Init(capacity uint) {
	h.data = make(map[interface{}]interface{}, capacity)
}

func (h *HashTable) Put(key interface{}, value interface{}) {
	h.data[key] = value
}

func (h *HashTable) Get(key interface{}) (interface{}, bool) {
	value, exists := h.data[key]
	return value, exists
}

func (h *HashTable) Delete(key interface{}) {
	delete(h.data, key)
}

func (h *HashTable) Keys() []interface{} {
	keys := make([]interface{}, 0, h.Size())
	for k := range h.data {
		keys = append(keys, k)
	}
	return keys
}
