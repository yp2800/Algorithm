package leetcode

// 左右乘积列表
func productExceptSelf(nums []int) []int {
	n := len(nums)
	//L和R分别表示左右两侧的乘积列表
	L, R, answer := make([]int, n), make([]int, n), make([]int, n)
	// L[i] 为索引左侧所有元素的乘积
	// 对于索引为 '0' 的元素，因为左侧没有元素，所以L[0]=1
	L[0] = 1
	for i := 1; i < n; i++ {
		L[i] = nums[i-1] * L[i-1]
	}
	// R[i] 为索引 i 右侧所有元素的乘积列表
	// 对于索引为 n-1 的元素，因为右侧没有元素，所以 R[n-1]=1
	R[n-1] = 1
	for i := n - 2; i >= 0; i-- {
		R[i] = nums[i+1] * R[i+1]
	}
	// 对于索引 i，除 nums[i] 之外其余各元素的乘积就是左侧所有元素的乘积乘以右侧所有元素的乘积
	for i := 0; i < n; i++ {
		answer[i] = L[i] * R[i]
	}
	return answer
}

func productExceptSelf2(nums []int) []int {
	n := len(nums)
	answer := make([]int, n)
	answer[0] = 1
	for i := 1; i < n; i++ {
		answer[i] = nums[i-1] * answer[i-1]
	}
	R := 1
	for i := n - 1; i >= 0; i-- {
		answer[i] = answer[i] * R
		R *= nums[i]
	}
	return answer
}
