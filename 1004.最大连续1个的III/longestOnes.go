package leetcode

import "sort"

// 二分查找
func longestOnes(nums []int, k int) int {
	n := len(nums)
	p := make([]int, n+1)
	ans := 0
	for i, num := range nums {
		p[i+1] = p[i] + 1 - num
	}
	for right, v := range p {
		left := sort.SearchInts(p, v-k)
		ans = max(ans, right-left)
	}
	return ans
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func longestOnes2(nums []int, k int) int {
	var left, lsum, rsum, ans int
	for right, num := range nums {
		rsum += 1 - num
		for lsum < rsum-k {
			lsum += 1 - nums[left]
			left++
		}
		ans = max(ans, right-left+1)
	}
	return ans
}
