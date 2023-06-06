package leetcode

func mergeAlternately(word1 string, word2 string) string {
	ans := ""
	n1 := len(word1)
	n2 := len(word2)
	for i, j := 0, 0; i < n1 || j < n2; i, j = i+1, j+1 {
		if i < n1 && j >= n2 {
			ans += word1[i:]
			break
		}
		if i >= n1 && j < n2 {
			ans += word2[j:]
			break
		}
		if i == n1 && j == n2 {
			break
		}
		ans += string(word1[i]) + string(word2[j])
	}
	return ans
}

func mergeAlternately2(word1 string, word2 string) string {
	n, m := len(word1), len(word2)
	ans := make([]byte, 0, n+m)
	for i := 0; i < n || i < m; i++ {
		if i < n {
			ans = append(ans, word1[i])
		}
		if i < m {
			ans = append(ans, word2[i])
		}
	}
	return string(ans)
}
