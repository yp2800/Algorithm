package leetcode

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	point := 0
	for i := 1; i < len(nums); i++ {
		if nums[point] != nums[i] {
			nums[point+1] = nums[i]
			point++
		}
	}
	return point + 1
}

func removeDuplicates2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	for i := len(nums) - 1; i > 0; i-- {
		if nums[i] == nums[i-1] {
			nums = append(nums[:i],nums[i+1:]...)
		}
	}
	return len(nums)
}
