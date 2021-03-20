package leetcode

import (
	"fmt"
	"strings"
)

func readBinaryWatch(num int) []string {
	result := []string{}
	for i := 0; i < 12; i++ {
		for j := 0; j < 60; j++ {
			b1 := fmt.Sprintf("%b", i)
			b2 := fmt.Sprintf("%b", j)
			sumOne := strings.Count(b1, "1") + strings.Count(b2, "1")
			if sumOne == num {
				result = append(result, fmt.Sprintf("%d:%02d", i, j))
			}
		}
	}
	return result
}

func readBinaryWatch2(num int) []string {
	res := []string{}
	backtrace(0, 10, num, []int{}, &res)
	return res
}

func backtrace(start, cap, target int, path []int, res *[]string) {
	if len(path) == target {
		min, hour := 0, 0
		for _, v := range path {
			if v >= 4 {
				min += 1 << (v - 4)
			} else {
				hour += 1 << v
			}
		}
		if min < 60 && hour < 12 {
			*res = append(*res, fmt.Sprintf("%d:%02d", hour, min))
		}
		path = []int{}
		return
	}
	for i := start; i < cap; i++ {
		path = append(path, i)
		backtrace(i+1, cap, target, path, res)
		path = path[:len(path)-1]
	}
}
