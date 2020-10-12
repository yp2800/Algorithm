package leetcode

import "sort"

func containsDuplicate(nums []int) bool {
	sort.Ints(nums)
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			return true
		}
	}
	return false
}

func containsDuplicate2(nums []int) bool {
	m:=make(map[int]int)
	for _, num := range nums {
		if _,ok := m[num];ok {
			return true
		} else {
			m[num] = 1
		}
	}
	return false
}