package exercises

func permute(nums []int) [][]int {
	ret := make([][]int, 0)
	usedNum := make(map[int]bool)
	dfs046(nums, usedNum, nil, &ret)
	return ret
}

func dfs046(nums []int, usedNums map[int]bool, path []int, ret *[][]int) {
	if len(path) == len(nums) {
		temp := make([]int, len(path))
		copy(temp, path)
		*ret = append(*ret, temp)
	}
	for _, curNum := range nums {
		if usedNums[curNum] {
			continue
		} // if>>

		usedNums[curNum] = true
		path = append(path, curNum)
		dfs046(nums, usedNums, path, ret)
		path = path[:len(path)-1]
		usedNums[curNum] = false
	} // for>
}
