package gokatas

import "sort"

func heightChecker(heights []int) int {
	unexpectedHeight := 0
	sortedArr := make([]int, len(heights))
	copy(sortedArr, heights)

	sort.Ints(sortedArr)

	for i := 0; i < len(sortedArr); i++ {
		if heights[i] != sortedArr[i] {
			unexpectedHeight++
		}
	}

	return unexpectedHeight
}
