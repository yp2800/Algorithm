package leetcode

func reverseString(s []byte) {
	for i, j := 0, len(s); i < j; i++ {
		s[i],s[j] = s[j],s[i]
		i++
		j--
	}
}
