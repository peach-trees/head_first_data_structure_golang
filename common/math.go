package common

import "math"

func MaxInt(values ...int) int {
	ret := math.MinInt32
	for i := range values {
		if values[i] > ret {
			ret = values[i]
		} // if>
	} // for>
	return ret
}
