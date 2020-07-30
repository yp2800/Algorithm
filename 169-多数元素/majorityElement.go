package leetcode

func majorityElement(nums []int) int {
	count := 0
	candidate := 0
	for _, value := range nums {
		if count == 0 {
			candidate = value
		}
		if value == candidate {
			count++
		} else {
			count--
		}
	}
	return candidate
}
