package gokatas

func solution(str string) []string {
	var result []string
	var split string

	for i := 0; i < len([]rune(str)); i++ {
		split = split + string(str[i])
		if len(split) == 2 {
			result = append(result, split)
			split = ""
		}
	}

	if len(split) == 1 {
		split = split + "_"
		result = append(result, split)
	}

	return result
}
