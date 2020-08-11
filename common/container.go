package common

// Container 所有的数据结构都是 Container 容器接口的实现
type Container interface {
	Size() int
	IsEmpty() bool
	Reset()
	Values() []interface{}
	String() string
}
