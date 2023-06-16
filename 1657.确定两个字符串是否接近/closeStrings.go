package leetcode

func closeStrings(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	} else {
		hash1 := make(map[rune]int)
		hash2 := make(map[rune]int)
		for _, v1 := range word1 {
			hash1[v1]++
		}
		for _, v2 := range word2 {
			hash2[v2]++
		}
		if len(hash1) != len(hash2) {
			return false
		}

		for k, _ := range hash1 {
			if _, ok := hash2[k]; !ok {
				return false
			}
		}
		for key, _ := range hash1 {
			for key2, _ := range hash2 {
				if hash1[key] == hash2[key2] {
					hash2[key2] = 0
					break
				}
			}
		}

		for _, v := range hash2 {
			if v != 0 {
				return false
			}
		}
		return true
	}

}
