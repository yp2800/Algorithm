package leetcode

func maxVowels(s string, k int) int {
	n := len(s)
	sum := 0
	for i := 0; i < k; i++ {
		if isVowel(s[i]) {
			sum++
		}
	}
	maxCount := sum
	for i := k; i < n; i++ {
		if isVowel(s[i-k]) {
			sum--
		}
		if isVowel(s[i]) {
			sum++
		}
		maxCount = max(maxCount, sum)
	}
	return maxCount
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func isVowel(ch byte) bool {
	switch ch {
	case 'a', 'e', 'i', 'o', 'u':
		return true
	}
	return false
}
