package leetcode

func titleToNumber(s string) int {
	result := 0
	for i := 0; i < len(s); i++ {
		tmp := s[i] - 'A' + 1
		result = result * 26 + int(tmp)
	}
	return result
}
