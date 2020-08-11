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
