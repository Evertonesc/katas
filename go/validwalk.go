package katas

func IsValidWalk(walk []rune) bool {
	if len(walk) != 10 {
		return false
	}

	xaxis, yaxis := 0, 0

	for i := 0; i < len(walk); i++ {
		if isXAxis(walk[i]) {
			if walk[i] == 'n' {
				xaxis++
			} else {
				xaxis--
			}

			continue
		}

		if walk[i] == 'w' {
			yaxis--
		} else {
			yaxis++
		}
	}

	if xaxis == 0 && yaxis == 0 {
		return true
	}

	return false
}

func isXAxis(r rune) bool {
	if r == 'n' || r == 's' {
		return true
	}

	return false
}
