package gokatas

func getWinner(arr []int, k int) int {
	winDb := make(map[int]int)

	if k >= len(arr) {
		max := arr[0]
		for _, v := range arr {
			if v > max {
				max = v
			}
		}

		return max
	}

	winner := arr[0]
	for i := 1; i < len(arr); i++ {
		if winner > arr[i] {
			winDb[winner]++
		} else {
			winner = arr[i]
			winDb[winner]++
		}

		if winDb[winner] == k {
			return winner
		}
	}

	return winner
}
