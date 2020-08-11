package hash_map

import (
	"bytes"
	"fmt"
)

func (h *HashMap) Size() int {
	return len(h.data)
}

func (h *HashMap) IsEmpty() bool {
	return len(h.data) == 0
}

func (h *HashMap) Reset() {
	h.data = make(map[interface{}]interface{})
}

func (h *HashMap) Values() []interface{} {
	values := make([]interface{}, 0, h.Size())
	for _, v := range h.data {
		values = append(values, v)
	} // for>
	return values
}

func (h *HashMap) String() string {
	bufStr := bytes.NewBufferString("[HashMap]")
	bufStr.WriteString(fmt.Sprintf("%+v", h.data))
	return bufStr.String()
}
