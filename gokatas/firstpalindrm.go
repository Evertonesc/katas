package gokatas

func firstPalindrome(words []string) string {
	for _, word := range words {
		left, right := 0, len(word)-1

		for left < len(word) {
			if right == 0 {
				return word
			}

			if word[left] != word[right] {
				break
			}

			left++
			right--
		}

	}

	return ""
}
