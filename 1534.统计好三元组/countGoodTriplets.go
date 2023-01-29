package leetcode

import "math"

func countGoodTriplets(arr []int, a int, b int, c int) int {
	cnt := 0
	n := len(arr)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				if math.Abs(float64(arr[i]-arr[j])) <= float64(a) && math.Abs(float64(arr[j]-arr[k])) <= float64(b) && math.Abs(float64(arr[i]-arr[k])) <= float64(c) {
					cnt++
				}
			}
		}
	}
	return cnt
}
