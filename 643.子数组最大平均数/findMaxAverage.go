package leetcode

func findMaxAverage(nums []int, k int) float64 {
	sum := 0
	for _, num := range nums[:k] {
		sum += num
	}
	maxSum := sum
	for i := k; i < len(nums); i++ {
		sum = sum - nums[i-k] + nums[i]
		maxSum = max(maxSum, sum)
	}
	return float64(maxSum) / float64(k)
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
