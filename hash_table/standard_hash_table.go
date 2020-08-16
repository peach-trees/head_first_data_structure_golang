package hash_table

type StandardHashTable struct {
	data map[interface{}]interface{}
}

func NewStandardHashTable() *StandardHashTable {
	return &StandardHashTable{data: make(map[interface{}]interface{})}
}

func (s *StandardHashTable) Put(key interface{}, value interface{}) {
	s.data[key] = value
}

func (s *StandardHashTable) Get(key interface{}) (interface{}, bool) {
	value, found := s.data[key]
	return value, found
}

func (s *StandardHashTable) Delete(key interface{}) {
	delete(s.data, key)
}

func (s *StandardHashTable) Keys() []interface{} {
	keys := make([]interface{}, 0, s.Size())
	for k := range s.data {
		keys = append(keys, k)
	}
	return keys
}

func (s *StandardHashTable) Values() []interface{} {
	values := make([]interface{}, 0, s.Size())
	for _, v := range s.data {
		values = append(values, v)
	}
	return values
}

func (s *StandardHashTable) IfEmpty() bool {
	return len(s.data) == 0
}

func (s *StandardHashTable) Size() int {
	return len(s.data)
}

func (s *StandardHashTable) Reset() {
	s.data = make(map[interface{}]interface{})
}
