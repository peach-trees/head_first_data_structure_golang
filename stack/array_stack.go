package stack

import (
	"bytes"
	"fmt"
)

type ArrayStack struct {
	data []interface{}
}

func NewArrayStack() *ArrayStack {
	return &ArrayStack{data: make([]interface{}, 0)}
}

func (a *ArrayStack) Push(value interface{}) {
	a.data = append(a.data, value)
}

func (a *ArrayStack) Pop() (interface{}, bool) {
	if len(a.data) == 0 {
		return nil, false
	}
	ret := a.data[len(a.data)-1]
	a.data = a.data[:len(a.data)-1]
	return ret, true
}

func (a *ArrayStack) Peek() (interface{}, bool) {
	if len(a.data) == 0 {
		return nil, false
	}
	return a.data[len(a.data)-1], true
}

func (a *ArrayStack) IfEmpty() bool {
	return len(a.data) == 0
}

func (a *ArrayStack) Size() int {
	return len(a.data)
}

func (a *ArrayStack) Reset() {
	a.data = make([]interface{}, 0)
}

func (a *ArrayStack) String() string {
	retBytes := bytes.NewBufferString("[Stack-top]")
	for i := len(a.data) - 1; i >= 0; i-- {
		retBytes.WriteString(fmt.Sprintf("->%+v", a.data[i]))
	}
	retBytes.WriteString("->[Stack-bottom]")
	return retBytes.String()
}
