package leetcode

func longestSubarray(nums []int) int {
	n := len(nums)
	pre, suf := make([]int, n), make([]int, n)
	pre[0] = nums[0]
	for i := 1; i < n; i++ {
		if nums[i] == 1 {
			pre[i] = pre[i-1] + 1
		} else {
			pre[i] = 0
		}
	}
	suf[n-1] = nums[n-1]
	for i := n - 2; i >= 0; i-- {
		if nums[i] == 1 {
			suf[i] = suf[i+1] + 1
		} else {
			suf[i] = 0
		}
	}

	ans := 0
	for i := 0; i < n; i++ {
		preSum := 0
		sufSum := 0
		if i == 0 {
			preSum = 0
		} else {
			preSum = pre[i-1]
		}
		if i == n-1 {
			sufSum = 0
		} else {
			sufSum = suf[i+1]
		}
		ans = max(ans, preSum+sufSum)
	}
	return ans
}
func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
