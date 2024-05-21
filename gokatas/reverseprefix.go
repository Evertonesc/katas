package gokatas

func reversePrefix(word string, ch byte) string {
	r := []rune(word)

	for i := 0; i < len(r); i++ {
		if byte(r[i]) == ch {
			rp := r[:i+1]
			sp := r[i+1:]

			for i, j := 0, len(rp)-1; i < j; i, j = i+1, j-1 {
				rp[i], rp[j] = rp[j], rp[i]
			}

			return string(rp) + string(sp)
		}
	}

	return word
}
