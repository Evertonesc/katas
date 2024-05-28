package gokatas

func maxWealth(accounts [][]int) int {
	maxWealth := 0
	for i := 0; i < len(accounts); i++ {
		rowVal := 0
		for j := 0; j < len(accounts[i]); j++ {
			rowVal += accounts[i][j]
		}

		if rowVal > maxWealth {
			maxWealth = rowVal
		}
	}

	return maxWealth
}
