package gokatas

// leetcode #1128
func numEquivDominoPairs(dominoes [][]int) int {
	pairs := 0
	m := make(map[int]int)

	for i := 0; i < len(dominoes); i++ {

		if dominoes[i][0] > dominoes[i][1] {
			dominoes[i][0], dominoes[i][1] = dominoes[i][1], dominoes[i][0]
		}

		dom := dominoes[i][0] + dominoes[i][1]

		if v, ok := m[dom]; ok {
			pairs += v
		}

		m[dom]++
	}

	return pairs
}
