package gokatas

func countDivisiblePairs(nums []int, k int) int {
	left, right := 0, 1
	last := len(nums) - 1
	pairs := 0

	if len(nums) == 1 {
		return 0
	}

	for left < right {
		if nums[left] == nums[right] {
			p := left * right
			if p%k == 0 {
				pairs++
			}
		}

		if right == last {
			left++
			if left == last {
				break
			}

			right = left + 1
		} else {
			right++
		}
	}

	return pairs
}
