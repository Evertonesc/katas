package gokatas

func findMatrix(nums []int) [][]int {
	var ans [][]int
	nmap := make(map[int]int)

	for len(nums) >= 1 {
		var subArr []int

		for i := 0; i < len(nums); i++ {
			_, ok := nmap[nums[i]]
			if !ok {
				nmap[nums[i]]++
			} else {
				continue
			}

			subArr = append(subArr, nums[i])
			nums = append(nums[:i], nums[i+1:]...)

			i--
		}

		ans = append(ans, subArr)
		nmap = make(map[int]int)
	}

	return ans
}
