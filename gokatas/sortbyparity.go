package gokatas

func sortArrayByParity(nums []int) []int {
	k := 0

	for i := 0; i < len(nums); i++ {
		if nums[i]%2 == 0 {
			nums[k], nums[i] = nums[i], nums[k]
			k++
		}
	}

	return nums
}
