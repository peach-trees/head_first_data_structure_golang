package exercises

import "sort"

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return nil
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	merged := make([][]int, 0)
	merged = append(merged, intervals[0])
	for i := 1; i < len(intervals); i++ {
		cur := intervals[i]
		last := merged[len(merged)-1]
		if cur[0] > last[1] {
			merged = append(merged, cur)
		} else {
			if cur[1] > last[1] {
				last[1] = cur[1]
			} // if>>>
		} // else>>
	} // for>
	return merged
}
