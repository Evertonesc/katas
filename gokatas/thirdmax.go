package gokatas

import "sort"

// search in the array
// if length less than 3, return the last item?
// when dealing with repeated numbers, use a map to track repetitions and etc...
//func thirdMax(nums []int) int {
//	var max1, max2, max3 *int
//
//	for _, num := range nums {
//		if max1 == nil || num > *max1 {
//			max3 = max2
//			max2 = max1
//			max1 = &num
//		} else if *max1 != num && (max2 == nil || num > *max2) {
//			max3 = max2
//			max2 = &num
//		} else if *max1 != num && *max2 != num && (max3 == nil || num > *max3) {
//			max3 = &num
//		}
//	}
//
//	if max3 == nil {
//		return *max1
//	}
//
//	return *max3
//}

func thirdMax(nums []int) int {
	seen := make(map[int]bool)
	result := make([]int, 0)

	for _, v := range nums {
		if !seen[v] {
			seen[v] = true
			result = append(result, v)
		}
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i] > result[j]
	})

	if len(result) >= 3 {
		return result[2]
	}

	return result[0]
}
