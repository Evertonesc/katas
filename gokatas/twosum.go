package gokatas

func twoSum(nums []int, target int) []int {
	sumMap := make(map[int]int)

	for i, v := range nums {
		r := target - v
		index, ok := sumMap[r]
		if ok {
			return []int{index, i}
		}

		sumMap[v] = i
	}

	return []int{}
}
