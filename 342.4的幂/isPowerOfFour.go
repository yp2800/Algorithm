package leetcode

import (
	"strconv"
	"strings"
)

func isPowerOfFour(n int) bool {
	if n < 1 {
		return false
	}

	s := strconv.FormatInt(int64(n), 4)
	return s[0:1] == "1" && strings.Count(s, "0") == len(s)-1
}

func isPowerOfFour2(n int) bool {
	return n > 0 && n&(n-1) == 0 && 0xaaaaaaaa&n == 0
}
