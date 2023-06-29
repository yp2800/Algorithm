package leetcode

import "sort"

// 排序
func merge(intervals [][]int) [][]int {
	n := len(intervals)
	if n == 0 {
		return nil
	}
	sort.SliceStable(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	ans := make([][]int, 0)
	for i := 0; i < n; i++ {
		L := intervals[i][0]
		R := intervals[i][1]
		length := len(ans)
		if length == 0 || ans[length-1][1] < L {
			ans = append(ans, []int{L, R})
		} else {
			ans[length-1][1] = max(ans[length-1][1], R)
		}
	}
	return ans
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
