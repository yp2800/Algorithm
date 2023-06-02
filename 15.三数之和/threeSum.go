package leetcode

import "sort"

// 排序+双指针
func threeSum(nums []int) [][]int {
	n := len(nums)
	sort.Ints(nums)
	ans := make([][]int, 0)
	// 遍历 nums
	for first := 0; first < n; first++ {
		// 需要和上一次枚举的数不相同
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		// third 对应的指针初始指向数组的最右端
		third := n - 1
		target := -1 * nums[first]
		//枚举 nums
		for second := first + 1; second < n; second++ {
			//需要和上次枚举的数不相同
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}
			//需要保证second的指针在third的左侧
			for second < third && nums[second]+nums[third] > target {
				third--
			}
			//	如果指针重合，随着second后续的增加，就不会有first + second + third = 0 并且second < third了，可以退循环
			if second == third {
				break
			}
			if nums[second]+nums[third] == target {
				ans = append(ans, []int{nums[first], nums[second], nums[third]})
			}
		}
	}
	return ans
}
