package leetcode

func getRow(rowIndex int) []int {
	if rowIndex == 0 {
		return []int{1}
	}

	res := [][]int{{1}}
	for i := 1; i <= rowIndex; i++ {
		m := []int{0}
		m = append(m, res[i-1]...)
		for j := 0; j < len(m)-1; j++ {
			m[j] = m[j] + m[j+1]
		}
		res = append(res, m)
	}
	return res[rowIndex]
}

func getRow2(rowIndex int) []int {
	nums := []int{1}
	for i := 1; i <= rowIndex; i++ {
		nums = append(nums, 1)
		for j := i - 1; j > 0; j-- {
			nums[j] += nums[j-1]
		}
	}
	return nums
}
