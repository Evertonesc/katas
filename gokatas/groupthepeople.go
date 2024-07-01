package gokatas

func groupThePeople(groupSizes []int) [][]int {
	var ans [][]int

	hm := make(map[int][]int)
	for i, v := range groupSizes {
		arr, ok := hm[groupSizes[i]]
		if !ok {
			hm[groupSizes[i]] = make([]int, 0)
			hm[groupSizes[i]] = append(hm[groupSizes[i]], i)

			continue
		}

		if len(arr) == v {
			ans = append(ans, arr)
			hm[groupSizes[i]] = []int{i}
		} else {
			arr = append(arr, i)
			hm[groupSizes[i]] = arr
		}
	}

	for _, v := range hm {
		ans = append(ans, v)
	}

	return ans
}
