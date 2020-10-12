package leetcode

func isIsomorphic(s string, t string) bool {
	return isIsomorphicHelper(s, t) && isIsomorphicHelper(t, s)
}

func isIsomorphicHelper(s, t string) bool {
	if len(s) != len(t) {
		return false
	}

	hm := make(map[rune]rune)

	for k, v := range s {
		if val, ok := hm[v]; ok {
			if rune(t[k]) != val {
				return false
			}
		} else {
			hm[v] = rune(t[k])
		}
	}
	return true
}

func isIsomorphic2(s, t string) bool {
	if len(s) != len(t) {
		return false
	}

	hm := make(map[rune]rune)
	hm_t := make(map[rune]rune)

	for k, v := range s {
		if val, ok := hm[v]; ok {
			if rune(t[k]) != val {
				return false
			}
		} else {
			hm[v] = rune(t[k])
		}

		if val, ok := hm_t[rune(t[k])]; ok {
			if val != v {
				return false
			}
		} else {
			hm_t[rune(t[k])] = v
		}
	}
	return true
}
