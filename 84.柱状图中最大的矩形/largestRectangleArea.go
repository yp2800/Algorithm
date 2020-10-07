package leetcode

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

// 暴力法 时间复杂度 O(n^3) 空间复杂度 O(1)
func largestRectangleArea(heights []int) int {
	maxArea := 0
	for i := 0; i < len(heights); i++ {
		for j := i; j < len(heights); j++ {
			minHigh := int(^uint(0) >> 1)
			for k := i; k <= j; k++ {
				minHigh = min(minHigh, heights[k])
				maxArea = max(maxArea, minHigh*(k-i+1))
			}
		}
	}
	return maxArea
}

// 优化的暴力
func largestRectangleArea2(heights []int) int {
	maxArea := 0
	for i := 0; i < len(heights); i++ {
		minHigh := int(^uint(0) >> 1)
		for j := i; j < len(heights); j++ {
			minHigh = min(minHigh, heights[j])
			maxArea = max(maxArea, minHigh*(j-i+1))
		}
	}
	return maxArea
}

// 分治法
func largestRectangleArea3(heights []int) int {
	return calculateArea(heights, 0, len(heights)-1)
}
func calculateArea(heights []int, start, end int) int {
	if start > end {
		return 0
	}
	minIndex := start
	for i := start; i <= end; i++ {
		if heights[minIndex] > heights[i] {
			minIndex = i
		}
	}
	return max(heights[minIndex]*(end-start+1), max(calculateArea(heights, start, minIndex-1), calculateArea(heights, minIndex+1, end)))
}

// todo 优化的分治
func largestRectangleArea4(heights []int) int {
	return heights[0]
}                                  

// 栈                                              ´
func largestRectangleArea5(heights []int) int {
	stack := []int{-1}
	maxArea := 0
	for i := 0; i < len(heights); i++ {
		for stack[len(stack)-1] != -1 && heights[stack[len(stack)-1]] >= heights[i] {
			pop := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			maxArea = max(maxArea, heights[pop]*(i-stack[len(stack)-1]-1))
		}
		stack = append(stack, i)
	}
	for stack[len(stack)-1] != -1 {
		pop := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		maxArea = max(maxArea, heights[pop]*(len(heights)-stack[len(stack)-1]-1))
	}
	return maxArea
}
