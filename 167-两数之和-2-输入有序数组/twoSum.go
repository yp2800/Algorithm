package leetcode

// 暴力
func twoSum(numbers []int, target int) []int {
	for i := range numbers {
		if numbers[i] <= target {
			for j := i + 1; j < len(numbers); j++ {
				if target-numbers[i] == numbers[j] {
					return []int{i + 1, j + 1}
				} else if target-numbers[i] < numbers[j] {
					break
				}
			}
		} else {
			return nil
		}
	}
	return []int{-1, -1}
}

// 二分查找
func twoSum2(numbers []int, target int) []int {
	for i := 0; i < len(numbers); i++ {
		low, high := i+1, len(numbers)-1
		for low <= high {
			mid := (high-low)/2 + low
			if numbers[mid] == target-numbers[i] {
				return []int{i + 1, mid + 1}
			} else if numbers[mid] > target-numbers[i] {
				high = mid - 1
			} else {
				low = mid + 1
			}
		}
	}
	return []int{-1, -1}
}

// 双指针法
func twoSum3(numbers []int, target int) []int {
	low, high := 0, len(numbers)-1
	for low < high {
		sum := numbers[low] + numbers[high]
		if sum == target {
			return []int{low + 1, high + 1}
		} else if sum > target {
			high--
		} else {
			low++
		}
	}
	return []int{-1, -1}
}
