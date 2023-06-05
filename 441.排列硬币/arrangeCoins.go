package leetocde

import (
	"math"
	"sort"
)

func arrangeCoins(n int) int {
	return sort.Search(n, func(k int) bool {
		k++
		return k*(k+1) > 2*n
	})
}

func arrangeCoins2(n int) int {
	return int((math.Sqrt(float64(8*n+1)) - 1) / 2)
}
