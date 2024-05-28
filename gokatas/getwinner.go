package gokatas

func getWinner(arr []int, k int) int {
	winDb := make(map[int]int)

	for {
		larger, smaller := 0, 0
		if arr[0] > arr[1] {
			larger, smaller = arr[0], arr[1]
			arr = append(arr[:1], arr[2:]...)
			arr = append(arr, smaller)

		} else {
			larger, smaller = arr[1], arr[0]
			arr = append(arr[1:])
			arr = append(arr, smaller)
		}

		if _, ok := winDb[larger]; !ok {
			winDb[larger]++
		} else {
			winDb[larger]++
		}
		v, _ := winDb[larger]
		if v == k {
			return larger
		}
	}
}
