package leetcode


func maxSubArray(nums []int) int {
return maxSubArrayHelper(nums, 0, len(nums)-1)
}

func maxSubArrayHelper(nums []int, left int, right int) int {
	if left == right {
		return nums[left]
	}
	mid := (left + right) / 2
	leftSum := maxSubArrayHelper(nums, left, mid)
	rightSum := maxSubArrayHelper(nums, mid+1, right)
	midSum := findMaxCrossingSubarray(nums, left, mid, right)
	maxSum := max(leftSum, rightSum)
	maxSum = max(maxSum, midSum)
	return maxSum
}

func findMaxCrossingSubarray(nums []int, left int, mid int, right int) int {
	INT_MAX := int(^uint(0) >> 1)
	INT_MIN := ^INT_MAX

	leftSum := INT_MIN
	sum := 0
	for i := mid; i >= left; i-- {
		sum += nums[i]
		leftSum = max(leftSum, sum)
	}

	rightSum := INT_MIN
	sum = 0
	for i := mid + 1; i <= right; i++ {
		sum += nums[i]
		rightSum = max(rightSum, sum)
	}

	return leftSum + rightSum
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
