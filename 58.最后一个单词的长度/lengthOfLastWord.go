package leetcode

import (
	"strings"
)

func lengthOfLastWord(s string) int {
	if len(s) == 0 {
		return 0
	}
	arr := strings.Split(strings.TrimSpace(s), " ")
	if len(arr) == 0 {
		return 0
	}
	return len(arr[len(arr)-1])
}

func lengthOfLastWord2(s string) int {
	idx := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != byte(' ') {
			idx++
		} else {
			if idx > 0 {
				return idx
			}
		}
	}
	return idx
}
