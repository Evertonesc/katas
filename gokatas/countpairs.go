package gokatas

// Another solution would be two for loops i and j
// where a inner loop for j starts in the position 1 in the array
func countPair(nums []int, target int) int {
	left, right := 0, 1
	result := 0

	for left < right {
		if right > len(nums)-1 {
			left++
			right = left + 1

			if right >= len(nums) {
				break
			}
		}

		if nums[left]+nums[right] < target {
			result++
		}

		right++
	}

	return result
}
