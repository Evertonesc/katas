package gokatas

import (
	"fmt"
	"sort"
)

// leetcode #15
func threeSum(nums []int) [][]int {
	triplets := [][]int{}
	m := make(map[string]int)
	n := len(nums)

	sort.Ints(nums)
	for i := 0; i < n-2; i++ {
		if i > 0 && nums[i+1] == nums[i-1] {
			continue
		}

		left, right := i+1, n-1
		for left < right {
			sum := nums[i] + nums[right] + nums[left]
			if sum == 0 {
				triplet := []int{nums[i], nums[right], nums[left]}
				tr := tripletRep(triplet)

				if _, ok := m[tr]; ok {
					break
				}

				triplets = append(triplets, triplet)
				m[tr]++

				for left < right && nums[left] == nums[left+1] {
					left++
				}

				for left < right && nums[right] == nums[right-1] {
					right--
				}

				left++
				right--

			} else if sum < 0 {
				left++
			} else {
				right--
			}
		}
	}

	return triplets
}

func tripletRep(triplet []int) string {
	return fmt.Sprintf("%d%d%d", triplet[0], triplet[1], triplet[2])
}
