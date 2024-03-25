package gokatas

// list b do not need to be modified
// 1 - check the edge cases i.e = arrays with different lengths

// for arrays with different sizes, you can do the following:
// 1 - get the smaller length to loop over
// 2 - create a map with the smaller array, store its values in the map and loop over the bigger array inserting
// values that are different from the ones inside the map.

func arrayDiff(a, b []int) []int {
	if len(a) == 0 {
		return []int{}
	}

	if len(b) == 0 {
		return a
	}

	var r []int
	bMap := make(map[int]bool)
	for i := 0; i < len(b); i++ {
		bMap[b[i]] = true
	}

	for i := 0; i < len(a); i++ {
		if !bMap[a[i]] {
			r = append(r, a[i])
		}
	}

	return r
}
