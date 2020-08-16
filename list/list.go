package list

type List interface {
	Append(values ...interface{})
	Prepend(values ...interface{})
	Get(index int) (interface{}, bool)
	Set(index int, value interface{})
	Delete(index int)
	Insert(index int, value interface{})
	IndexOf(value interface{}) int
	IfEmpty() bool
	Size() int
	Reset()
	String() string
}
