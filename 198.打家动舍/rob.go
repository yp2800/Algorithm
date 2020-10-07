package leetcode

func rob(nums []int) int {
	prevMax := 0
	currMax := 0
	for _, value := range nums {
		tmp := currMax
		currMax = max(prevMax+value, currMax)
		prevMax = tmp
	}
	return currMax
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
