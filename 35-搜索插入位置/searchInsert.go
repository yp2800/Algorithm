package leetcode

// 暴力解法
func searchInsert(nums []int, target int) int {
	for i, num := range nums {
		if num >= target {
			return i
		}
	}
	return len(nums)
}

//二分查找
func searchInsert2(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left+(right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	// 防溢出
	if left == len(nums) {
		return left
	}
	if target <= nums[left] {
		return left
	} else {
		return left + 1
	}
}
