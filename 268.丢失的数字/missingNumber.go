package leetcode

import (
	"sort"
)

// 高斯求和公式
func missingNumber(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return len(nums)*(len(nums)+1)/2 - sum
}

func missingNumber2(nums []int) int {
	missing := len(nums)
	for i := 0; i < len(nums); i++ {
		missing ^= i ^ nums[i]
	}
	return missing
}

func missingNumber3(nums []int) int {
	m := map[int]bool{}
	for _, num := range nums {
		m[num] = true
	}
	for i := 0; i <= len(nums); i++ {
		if _, ok := m[i]; !ok {
			return i
		}
	}
	return -1
}

func missingNumber4(nums []int) int {
	sort.Ints(nums)
	if nums[0] != 0 {
		return 0
	} else if nums[len(nums)-1] != len(nums) {
		return len(nums)
	}

	for i := 1; i < len(nums); i++ {
		expectedNum := nums[i-1] + 1
		if expectedNum != nums[i] {
			return expectedNum
		}
	}
	return -1
}
