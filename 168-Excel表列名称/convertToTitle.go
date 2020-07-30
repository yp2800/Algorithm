package leetcode

func convertToTitle(n int) string {
	result := ""
	for n > 0 {
		m:=n%26
		if m == 0 {
			m=26
			n--
		}
		n = n / 26
		result = string(m+64)+result
	}
	return result
}
