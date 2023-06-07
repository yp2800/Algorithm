package leetcode

// 枚举
func gcdOfStrings(str1 string, str2 string) string {
	n1, n2 := len(str1), len(str2)
	minString := minStrings(str1, str2)
	for i := min(n1, n2); i >= 1; i-- {
		if n1%i == 0 && n2%i == 0 {
			x := minString[:i]
			if check(x, str1) && check(x, str2) {
				return x
			}
		}
	}
	return ""
}

func check(t, s string) bool {
	lenx := len(s) / len(t)
	ans := ""
	for i := 0; i < lenx; i++ {
		ans = ans + t
	}
	return ans == s
}
func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func minStrings(str1, str2 string) string {
	if len(str1) < len(str2) {
		return str1
	}
	return str2
}
