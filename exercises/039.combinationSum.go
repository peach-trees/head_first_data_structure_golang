package exercises

import "sort"

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	ret := make([][]int, 0)
	dfs039(candidates, target, 0, nil, &ret)
	return ret
}

func dfs039(candidates []int, target int, start int, path []int, ret *[][]int) {
	if target == 0 {
		temp := make([]int, len(path))
		copy(temp, path)
		*ret = append(*ret, temp)
		return
	} // if>
	for i := start; i < len(candidates); i++ {
		if candidates[start] > target {
			return
		}
		path = append(path, candidates[i])
		dfs039(candidates, target-candidates[i], i, path, ret)
		path = path[:len(path)-1]
	} // for>
}
