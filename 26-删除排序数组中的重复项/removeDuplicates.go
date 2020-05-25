package leetcode

func removeDuplicates(nums []int) int {
	point := 0
	for i:=1;i<len(nums);i++ {
		if nums[point] != nums[i] {
			nums[point+1] = nums[i]
			point++
		}
	}
	return point + 1
}