package leetcode

import "sort"

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	ans := 0
	m, n := len(g), len(s)
	for i, j := 0, 0; i < m && j < n; i++ {
		for j < n && g[i] > s[j] {
			j++
		}
		if j < n {
			ans++
			j++
		}
	}
	return ans
}
