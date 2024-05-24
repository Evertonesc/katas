package gokatas

// leetcode #832
func flipAndInvertImage(image [][]int) [][]int {
	for i := 0; i < len(image); i++ {
		last := len(image[i]) - 1
		for j := 0; j < len(image[i]); j++ {

			if j != last {
				image[i][j], image[i][last] = image[i][last], image[i][j]
				last--
			}

			if j == last {
				for k := 0; k < len(image[i]); k++ {
					if image[i][k] == 0 {
						image[i][k] = 1
					} else {
						image[i][k] = 0
					}
				}

				break
			}
		}
	}

	return image
}
