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

func maxOperations2(nums []int, k int) int {
	m := make(map[int]int)
	res := 0
	for _, num := range nums {
		if m[k-num] > 0 {
			m[k-num]--
			res++
		} else {
			m[num]++
		}
	}
	return res
}
