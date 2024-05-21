package gokatas

func arithmeticTriplets(nums []int, diff int) int {
	triplets := 0
	n := len(nums)

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				if nums[k]-nums[j] == diff && nums[j]-nums[i] == diff {
					triplets++
				}
			}
		}
	}

	return triplets
}
