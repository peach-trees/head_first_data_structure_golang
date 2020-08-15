package list

import (
	"bytes"
	"fmt"
)

type ArrayList struct {
	data []interface{}
	size int
}

func NewArrayList(values ...interface{}) *ArrayList {
	list := &ArrayList{size: 0, data: make([]interface{}, 0)}
	if len(values) > 0 {
		list.Append(values...)
	}
	return list
}

func (a *ArrayList) checkIndex(index int) bool {
	return index >= 0 && index < a.size
}

func (a *ArrayList) upScale(newDataLen int) {
	currentCapacity := cap(a.data)
	if a.size+newDataLen >= currentCapacity {
		newCapacity := (currentCapacity + newDataLen) << 1
		newData := make([]interface{}, newCapacity, newCapacity)
		copy(newData, a.data)
		a.data = newData
	} // if>
}

func (a *ArrayList) Append(values ...interface{}) {
	a.upScale(len(values))
	for i := range values {
		a.data[a.size] = values[i]
		a.size++
	} // for>
}

func (a *ArrayList) Prepend(values ...interface{}) {
	newValueLen := len(values)
	a.upScale(newValueLen)
	copy(a.data[newValueLen:], a.data[:a.size])
	copy(a.data[:newValueLen], values)
	a.size += newValueLen
}

func (a *ArrayList) Get(index int) (interface{}, bool) {
	if !a.checkIndex(index) {
		return nil, false
	}
	return a.data[index], true
}

func (a *ArrayList) Set(index int, value interface{}) {
	if !a.checkIndex(index) {
		return
	} // if>
	a.data[index] = value
}

func (a *ArrayList) Delete(index int) {
	if !a.checkIndex(index) {
		return
	}
	copy(a.data[index:], a.data[index+1:a.size])
	a.size--
}

func (a *ArrayList) Insert(index int, value interface{}) {
	if !a.checkIndex(index) {
		return
	} // if>
	a.upScale(1)
	for i := a.size - 1; i >= index; i-- {
		a.data[i+1] = a.data[i]
	} // for>
	a.data[index] = value
	a.size++
}

func (a *ArrayList) IndexOf(value interface{}) int {
	if a.size == 0 {
		return -1
	} // if>
	for index := range a.data {
		if a.data[index] == value {
			return index
		} // if>>
	} // for>
	return -1
}

func (a *ArrayList) IfEmpty() bool {
	return a.size == 0
}

func (a *ArrayList) Size() int {
	return a.size
}

func (a *ArrayList) Reset() {
	a.size = 0
	a.data = make([]interface{}, 0)
}

func (a *ArrayList) String() string {
	retBytes := bytes.NewBufferString("[ArrayList]")
	for i := 0; i < a.size; i++ {
		retBytes.WriteString(fmt.Sprintf("->%+v", a.data[i]))
	}
	return retBytes.String()
}
