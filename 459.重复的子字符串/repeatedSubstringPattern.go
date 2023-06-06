package leetcode

import "strings"

// 枚举
func repeatedSubstringPattern(s string) bool {
	n := len(s)
	for i := 1; i*2 <= n; i++ {
		if n%i == 0 {
			match := true
			for j := i; j < n; j++ {
				if s[j] != s[j-i] {
					match = false
					break
				}
			}
			if match {
				return true
			}
		}
	}
	return false
}

// 字符串匹配
func repeatedSubstringPattern2(s string) bool {
	s1 := s + s
	s2 := s1[1 : len(s1)-1]
	return strings.Contains(s2, s)
}
