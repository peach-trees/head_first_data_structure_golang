package exercises

import "sort"

func groupAnagrams(strs []string) [][]string {
	checkMap := make(map[string][]string)
	for _, v := range strs {
		if _, exists := checkMap[strKey(v)]; !exists {
			checkMap[strKey(v)] = make([]string, 0)
		} // if>>
		checkMap[strKey(v)] = append(checkMap[strKey(v)], v)
	} // for>

	ret := make([][]string, 0)
	for k := range checkMap {
		ret = append(ret, checkMap[k])
	}
	return ret
}

func strKey(str string) string {
	chArr := []byte(str)
	sort.Slice(chArr, func(i, j int) bool {
		return chArr[i] < chArr[j]
	})
	return string(chArr)
}
