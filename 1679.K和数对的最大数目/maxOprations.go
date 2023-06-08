package leetcode

import "sort"

func maxOperations(nums []int, k int) int {
	sort.Ints(nums)
	left, right := 0, len(nums)-1
	ansCount := 0
	for left < right {
		temp := nums[left] + nums[right]
		if temp < k {
			left++
		} else if temp > k {
			right--
		} else {
			ansCount++
			left++
			right--
		}
	}
	return ansCount
}
