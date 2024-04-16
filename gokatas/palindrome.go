package gokatas

import "strconv"

func isPalindrome(x int) bool {
	strInt := strconv.Itoa(x)
	last := len(strInt) - 1

	for i := 0; i < len([]rune(strInt)); i++ {
		if strInt[i] != strInt[last-i] {
			return false
		}
	}

	return true
}
