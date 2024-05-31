package gokatas

import (
	"math"
)

func shortestsToChar(s string, c byte) []int {
	var charSpots []int
	sr := []byte(s)

	for i, v := range sr {
		if v == c {
			charSpots = append(charSpots, i)
		}
	}

	answer := make([]int, len(s))
	j := 0
	for i := 0; i < len([]rune(s)); i++ {
		if i > charSpots[j] {
			cs := abs(i - charSpots[j])
			n := abs(i - charSpots[j+1])

			if cs > n {
				j++
			}
		}

		answer[i] = abs(i - charSpots[j])
	}

	return answer
}

func abs(n int) int {
	return int(math.Abs(float64(n)))
}
