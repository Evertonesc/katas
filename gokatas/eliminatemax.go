package gokatas

import "sort"

func eliminateMax(dist []int, speed []int) int {
	n := len(dist)
	t := make([]int, n)

	for i := 0; i < n; i++ {
		t[i] = (dist[i] - 1) / speed[i]
	}

	sort.Ints(t)

	for i := 0; i < len(t); i++ {
		if t[i] < i {
			return i
		}
	}

	return n
}
