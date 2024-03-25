package gokatas

func findDuplicate(nums []int) int {
	seen := make(map[int]bool)
	result := 0
	for i := 0; i < len(nums); i++ {
		if !seen[nums[i]] {
			seen[nums[i]] = true
		} else {
			result = nums[i]
			break
		}
	}
	return result
}
