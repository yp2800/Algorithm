package leetcode

func maxArea(height []int) int {
	l, r := 0, len(height)-1
	ans := 0
	for l < r {
		area := min(height[l], height[r]) * (r - l)
		ans = max(ans, area)
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return ans
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
