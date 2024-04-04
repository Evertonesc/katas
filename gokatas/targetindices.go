package gokatas

import "sort"

func targetIndices(nums []int, target int) []int {
	indices := make([]int, 0)

	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			indices = append(indices, i)
		}
	}

	return indices
}
