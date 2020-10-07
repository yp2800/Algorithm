package leetcode

import "strings"

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	INT_MAX := int(^uint(0) >> 1)
	minLen := INT_MAX
	for _, v := range strs {
		minLen = min(len(v), minLen)
	}
	low := 1
	high := minLen
	for low <= high {
		mid := (low + high) / 2
		if isCommonPrefix(strs, mid) {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return strs[0][:(low+high)/2]
}

func isCommonPrefix(strs []string, length int) bool {
	str1 := strs[0][:length]
	for i := 1; i < len(strs); i++ {
		if strings.Index(strs[i], str1) != 0 {
			return false
		}
	}
	return true
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
