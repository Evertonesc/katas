package gokatas

func removeElement(nums []int, val int) int {
	k := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[i], nums[k] = nums[k], nums[i]
			k++
		}
	}

	return k
}
