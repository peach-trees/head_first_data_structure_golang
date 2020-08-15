package queue

type Dequeue struct {
	head     int
	tail     int
	size     int
	capacity int
	data     []interface{}
}

func (d *Dequeue) Size() int {
	return d.size
}

func (d *Dequeue) IsEmpty() bool {
	return d.size == 0
}

func (d *Dequeue) PushBack(value interface{}) {
	d.upScale()
	d.data[d.tail] = value
	d.tail = d.next(d.tail)
	d.size++
}

func (d *Dequeue) PeekBack() interface{} {
	if d.size <= 0 {
		return nil
	}
	return d.data[d.head]
}

func (d *Dequeue) PopFront() interface{} {
	if d.size <= 0 {
		return nil
	}
	ret := d.data[d.head]
	d.data[d.head] = nil
	d.head = d.next(d.head)
	d.size--
	return ret
}

func (d *Dequeue) PushFront(value interface{}) {

}

func (d *Dequeue) PeekFront() interface{} {
	if d.size <= 0 {
		return nil
	}
	return d.data[d.prev(d.tail)]
}

// prev returns the previous buffer position wrapping around buffer.
func (d *Dequeue) prev(i int) int {
	return (i - 1) & (len(d.data) - 1) // bitwise modulus
}

// next returns the next buffer position wrapping around buffer.
func (d *Dequeue) next(i int) int {
	return (i + 1) & (len(d.data) - 1)
}

func (d *Dequeue) upScale() {
	if len(d.data) == 0 {
		if d.capacity == 0 {
			d.capacity = defaultCapacity
		} // if>>
		d.data = make([]interface{}, d.capacity)
		return
	} // if>
	if d.size == len(d.data) {
		d.resize()
	} // if>
}

func (d *Dequeue) resize() {
	newBuf := make([]interface{}, d.size<<1)
	if d.tail > d.head {
		copy(newBuf, d.data[d.head:d.tail])
	} else {
		n := copy(newBuf, d.data[d.head:])
		copy(newBuf[n:], d.data[:d.tail])
	} // else>
	d.head, d.tail, d.data = 0, d.size, newBuf
}
