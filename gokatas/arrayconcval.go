package gokatas

import (
	"fmt"
	"strconv"
)

func findTheArrayConcVal(nums []int) int64 {
	concVal := 0

	for len(nums) >= 1 {
		if len(nums) == 1 {
			concVal += nums[0]
			break
		}

		strVal := fmt.Sprintf("%d%d", nums[0], nums[len(nums)-1])

		intVal, _ := strconv.Atoi(strVal)

		concVal += intVal

		nums = nums[1 : len(nums)-1]
	}

	return int64(concVal)
}
