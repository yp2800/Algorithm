package leetcode

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	n := len(nums)
	sort.Ints(nums)
	best := math.MaxInt32
	//根据差值的绝对值来更新答案
	update := func(cur int) {
		if abs(cur-target) < abs(best-target) {
			best = cur
		}
	}
	//枚举 nums
	for first := 0; first < n; first++ {
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		//	使用双指针枚举 second 和 third
		second, third := first+1, n-1
		for second < third {
			sum := nums[first] + nums[second] + nums[third]
			if sum == target {
				return target
			}
			update(sum)
			if sum > target {
				third--
				for second < third && nums[third] == nums[third+1] {
					third--
				}
			} else {
				second++
				for second < third && nums[second] == nums[second-1] {
					second++
				}
			}
		}
	}
	return best
}

func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}
