package exercises

func twoSum(nums []int, target int) []int {
	if len(nums) <= 1 {
		return []int{-1, -1}
	}

	checkMap := make(map[int]int)
	for i := range nums {
		diff := target - nums[i]
		if index, exists := checkMap[diff]; exists {
			return []int{i, index}
		} else {
			checkMap[nums[i]] = i
		} // else>>
	} // for>
	return []int{-1, -1}
}
