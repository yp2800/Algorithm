package leetcode

import "fmt"

func equalPairs(grid [][]int) int {
	n := len(grid)
	m := make(map[string]int)
	for _, row := range grid {
		m[fmt.Sprint(row)]++
	}
	res := 0
	for j := 0; j < n; j++ {
		var arr []int
		for i := 0; i < n; i++ {
			arr = append(arr, grid[i][j])
		}
		if val, ok := m[fmt.Sprint(arr)]; ok {
			res += val
		}
	}
	return res
}
