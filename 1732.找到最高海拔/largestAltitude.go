package leetcode

func largestAltitude(gain []int) int {
	total := 0
	ans := 0
	for _, x := range gain {
		total += x
		ans = max(ans, total)
	}
	return ans
}

func max(x, y int) int {
	if x < y {
		return y
	}

	return x
}
