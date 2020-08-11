package common

import "sort"

// Comparator 比较器, 用于将 a 和 b 参数转换为相同的数据类型, 并返回比较结果.
// return -1, a < b
// return 0, a == b
// return 1, a > b
type Comparator func(a, b interface{}) int

// SortWrapper 将 values 包装为 可排序 sort-able 的结构
type SortWrapper struct {
	values []interface{}
	Comparator
}

func (s SortWrapper) Len() int {
	return len(s.values)
}

func (s SortWrapper) Swap(i, j int) {
	s.values[i], s.values[j] = s.values[j], s.values[i]
}

func (s SortWrapper) Less(i, j int) bool {
	return s.Comparator(s.values[i], s.values[j]) < 0
}

func IntComparator(a, b interface{}) int {
	aReflected := a.(int)
	bReflected := b.(int)
	switch {
	case aReflected > bReflected:
		return 1
	case aReflected < bReflected:
		return -1
	default:
		return 0
	}
}

func Int8Comparator(a, b interface{}) int {
	aReflected := a.(int8)
	bReflected := b.(int8)
	switch {
	case aReflected > bReflected:
		return 1
	case aReflected < bReflected:
		return -1
	default:
		return 0
	}
}

func Int32Comparator(a, b interface{}) int {
	aReflected := a.(int32)
	bReflected := b.(int32)
	switch {
	case aReflected > bReflected:
		return 1
	case aReflected < bReflected:
		return -1
	default:
		return 0
	}
}

func Int64Comparator(a, b interface{}) int {
	aReflected := a.(int64)
	bReflected := b.(int64)
	switch {
	case aReflected > bReflected:
		return 1
	case aReflected < bReflected:
		return -1
	default:
		return 0
	}
}

func UIntComparator(a, b interface{}) int {
	aReflected := a.(uint)
	bReflected := b.(uint)
	switch {
	case aReflected > bReflected:
		return 1
	case aReflected < bReflected:
		return -1
	default:
		return 0
	}
}

func UInt8Comparator(a, b interface{}) int {
	aReflected := a.(uint8)
	bReflected := b.(uint8)
	switch {
	case aReflected > bReflected:
		return 1
	case aReflected < bReflected:
		return -1
	default:
		return 0
	}
}

func UInt32Comparator(a, b interface{}) int {
	aReflected := a.(uint32)
	bReflected := b.(uint32)
	switch {
	case aReflected > bReflected:
		return 1
	case aReflected < bReflected:
		return -1
	default:
		return 0
	}
}

func UInt64Comparator(a, b interface{}) int {
	aReflected := a.(uint64)
	bReflected := b.(uint64)
	switch {
	case aReflected > bReflected:
		return 1
	case aReflected < bReflected:
		return -1
	default:
		return 0
	}
}

func Float64Comparator(a, b interface{}) int {
	aReflected := a.(float64)
	bReflected := b.(float64)
	switch {
	case aReflected > bReflected:
		return 1
	case aReflected < bReflected:
		return -1
	default:
		return 0
	}
}

// ByteComparator provides a basic comparison on byte
func ByteComparator(a, b interface{}) int {
	aReflected := a.(byte)
	bReflected := b.(byte)
	switch {
	case aReflected > bReflected:
		return 1
	case aReflected < bReflected:
		return -1
	default:
		return 0
	}
}

// SignedNumberComparator 将数字参数统一按照 int64 进行比较
var SignedNumberComparator = Int64Comparator

// UnSignedNumberComparator 将数字参数统一按照 uint64 进行比较
var UnSignedNumberComparator = UInt64Comparator

// Sort values by the comparator in place
func Sort(values []interface{}, comparator Comparator) {
	sort.Sort(SortWrapper{
		values:     values,
		Comparator: comparator,
	})
}
