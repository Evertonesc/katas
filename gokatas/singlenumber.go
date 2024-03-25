package gokatas

// no empty array of integers
// every element appears twice except for one.
func singleNumber(nums []int) int {
	result := 0
	for _, num := range nums {
		result ^= num
	}

	return result
}
