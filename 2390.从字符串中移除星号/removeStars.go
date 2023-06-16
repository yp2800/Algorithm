package leetcode

func removeStars(s string) string {
	answer := []rune{}
	for _, v := range s {
		if v == '*' {
			answer = answer[:len(answer)-1]
		} else {
			answer = append(answer, v)
		}
	}
	return string(answer)
}
