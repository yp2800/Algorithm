package leetcode

func twoSum(nums []int, target int) []int {
	result := make(map[int]int)
	for i, v := range nums {
		if k, ok := result[target-v]; ok {
			return []int{k, i}
		}
		result[v] = i
	}
	return nil
}
