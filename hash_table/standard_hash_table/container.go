package standard_hash_table

import (
	"bytes"
	"fmt"
)

func (h *HashTable) Size() int {
	return len(h.data)
}

func (h *HashTable) IsEmpty() bool {
	return len(h.data) == 0
}

func (h *HashTable) Reset() {
	h.data = make(map[interface{}]interface{})
}

func (h *HashTable) Values() []interface{} {
	values := make([]interface{}, 0, h.Size())
	for _, v := range h.data {
		values = append(values, v)
	} // for>
	return values
}

func (h *HashTable) String() string {
	bufStr := bytes.NewBufferString("[HashTable]")
	bufStr.WriteString(fmt.Sprintf("%+v", h.data))
	return bufStr.String()
}
