package gokatas

import (
	"fmt"
	"sort"
)

// the solution set cannot contian duplicate triplets
// all the triplets values must not be equal and its sums must be equal to 0

// leetcode #15
func threeSum(nums []int) [][]int {
	triplets := [][]int{}
	m := make(map[string]int)

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			for k := j + 1; k < len(nums); k++ {
				if nums[i]+nums[j]+nums[k] == 0 {
					triplet := []int{nums[i], nums[j], nums[k]}
					sort.Ints(triplet)
					tr := tripletRep(triplet)

					if _, ok := m[tr]; ok {
						continue
					}

					triplets = append(triplets, triplet)
					m[tr]++
				}
			}
		}
	}

	return triplets
}

func tripletRep(triplet []int) string {
	return fmt.Sprintf("%d%d%d", triplet[0], triplet[1], triplet[2])
}
