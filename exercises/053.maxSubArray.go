package exercises

func maxSubArray(nums []int) int {
	for i := 1; i < len(nums); i++ {
		if nums[i-1]+nums[i] > nums[i] {
			nums[i] += nums[i-1]
		} else {
			continue
		} // else>>
	} // for>

	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
		} // if>>
	} // for>
	return max
}
