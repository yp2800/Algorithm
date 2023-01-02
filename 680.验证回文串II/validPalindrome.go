package leetcode

func validPalindrome(s string) bool {
	for low, high := 0, len(s)-1; low < high; {
		if s[low] == s[high] {
			low++
			high--
			continue
		}
		flag1, flag2 := true, true
		for i, j := low, high-1; i < j; i, j = i+1, j-1 {
			if s[i] != s[j] {
				flag1 = false
				break
			}
		}
		for i, j := low+1, high; i < j; i, j = i+1, j-1 {
			if s[i] != s[j] {
				flag2 = false
				break
			}
		}
		return flag1 || flag2
	}
	return true
}
