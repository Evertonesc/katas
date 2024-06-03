package gokatas

// leetcode #1266
func minTimeToVisitAllPoints(points [][]int) int {
	minTime := 0

	for i := 0; i < len(points)-1; {
		if points[i][0] < points[i+1][0] {
			points[i][0]++
		} else if points[i][0] > points[i+1][0] {
			points[i][0]--
		}

		if points[i][1] < points[i+1][1] {
			points[i][1]++
		} else if points[i][1] > points[i+1][1] {
			points[i][1]--
		}

		if points[i][0] == points[i+1][0] && points[i][1] == points[i+1][1] {
			i++
		}

		minTime++
	}

	return minTime
}
