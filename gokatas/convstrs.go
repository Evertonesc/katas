package gokatas

import (
	"regexp"
	"strings"
	"unicode"
)

func ToCamelCase(s string) string {
	const pattern = `[-_]`

	words := regexp.MustCompile(pattern).Split(s, -1)

	var result string
	result = result + words[0]

	for i := 1; i < len(words); i++ {
		if words[i] == "" {
			continue
		}

		if !unicode.IsUpper([]rune(words[i])[0]) {
			sw := words[i][1:]
			r := strings.ToUpper(string(words[i][0]))
			cw := r + sw
			result = result + cw
		} else {
			result = result + words[i]
		}

	}

	return result
}
