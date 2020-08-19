package exercises

import (
	"head_first_data_structure_golang/common"
	"strings"
)

func lengthOfLongestSubstring(s string) int {
	if len(s) < 2 {
		return len(s)
	}
	var ret int
	var left int
	var tempStr string
	for i := 0; i < len(s); i++ {
		if strings.IndexByte(tempStr, s[i]) == -1 {
			tempStr = s[left : i+1]
			ret = common.MaxInt(ret, len(tempStr))
		} else {
			left += strings.IndexByte(s[left:i], s[i]) + 1
		} // else>>
	} // for>
	return ret
}
