package leetcode

import (
	"strconv"
	"strings"
)

func fizzBuzz(n int) []string {
	res := make([]string, 0)
	for i := 1; i < n+1; i++ {
		switch {
		case i%3 == 0 && i%5 == 0:
			res = append(res, "FizzBuzz")
		case i%3 == 0:
			res = append(res, "Fizz")
		case i%5 == 0:
			res = append(res, "Buzz")
		default:
			res = append(res, strconv.Itoa(i))
		}
	}
	return res
}

func fizzBuzz2(n int) []string {
	res := make([]string, 0)
	for i := 1; i < n+1; i++ {
		sb := &strings.Builder{}
		if i%3 == 0 {
			sb.WriteString("Fizz")
		}
		if i%5 == 0 {
			sb.WriteString("Buzz")
		}
		if sb.Len() == 0 {
			sb.WriteString(strconv.Itoa(i))
		}
		res = append(res, sb.String())
	}
	return res
}
