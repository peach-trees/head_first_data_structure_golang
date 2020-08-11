package hash_map

type HashMap struct {
	data map[interface{}]interface{}
}

func NewHashMap() *HashMap {
	return &HashMap{
		data: make(map[interface{}]interface{}),
	}
}

func (h *HashMap) Put(key interface{}, value interface{}) {
	h.data[key] = value
}

func (h *HashMap) Get(key interface{}) (interface{}, bool) {
	value, exists := h.data[key]
	return value, exists
}

func (h *HashMap) Delete(key interface{}) {
	delete(h.data, key)
}

func (h *HashMap) Keys() []interface{} {
	keys := make([]interface{}, 0, h.Size())
	for k := range h.data {
		keys = append(keys, k)
	}
	return keys
}
