package list

type ArrayList struct {
	data []interface{}
	size int
}

func (a *ArrayList) checkIndex(index int) bool {
	return index >= 0 && index < a.size
}

func (a *ArrayList) resize(capacity int) {
	newData := make([]interface{}, 0, capacity)
	copy(newData, a.data)
	a.data = newData
}

func (a *ArrayList) upScale(newDataLen int) {
	currentCapacity := cap(a.data)
	if a.size + newDataLen >= currentCapacity {
		newCapacity := (currentCapacity + newDataLen) << 1
		a.resize(newCapacity)
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

}

func (a *ArrayList) Delete(index int) {
	if !a.checkIndex(index) {
		return
	}
	a.data[index] = nil
	copy(a.data[index:], a.data[index+1:a.size])
	a.size--
}

func (a *ArrayList) Insert(index int, values ...interface{}) {

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

Sort(comparator common.Comparator)
Empty() bool
Size() int
Clear()
