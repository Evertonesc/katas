package gokatas

// leetcode #2744

func maxNumStrPairs(words []string) int {
	strPairs := 0

	for i := 0; i < len(words); i++ {
		for j := i + 1; j < len(words); j++ {
			if words[i] == reverseStr(words[j]) {
				strPairs++
			}
		}
	}

	return strPairs
}

func reverseStr(str string) (result string) {
	for _, v := range str {
		result = string(v) + result
	}

	return
}
