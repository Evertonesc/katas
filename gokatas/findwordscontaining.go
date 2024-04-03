package gokatas

// 0-indexed array
// array of indices representing the words that contain the character x
func findWordsContaining(words []string, x byte) []int {
	r := make([]int, 0)
	for i, word := range words {
		for _, w := range []byte(word) {
			if w == x {
				r = append(r, i)
				break
			}
		}
	}

	return r
}
