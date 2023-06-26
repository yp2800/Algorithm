package leetcode

// 二分查找
func search(nums []int, target int) int {
	n := len(nums)
	if n == 0 {
		return -1
	}
	l, r := 0, n-1
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] == target {
			return mid
		}
		//左侧子数组有序
		if nums[0] <= nums[mid] {
			if nums[0] <= target && target < nums[mid] {
				r = mid - 1
			} else {

				l = mid + 1
			}
		} else {
			//	右侧子数组有序
			if nums[mid] < target && target <= nums[n-1] {
				l = mid + 1
			} else {
				//右侧无序
				r = mid - 1
			}
		}
	}
	return -1
}
