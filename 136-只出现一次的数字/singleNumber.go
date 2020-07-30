package leetcode

func singleNumber(nums []int) int {
	a := 0
	for _, value := range nums {
		a ^= value
	}
	return a
}
