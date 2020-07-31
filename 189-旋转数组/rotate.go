package leetcode

import (
	"fmt"
)

func rotate(nums []int, k int) {
	var temp, previous int
	for i := 0; i < k; i++ {
		previous = nums[len(nums)-1]
		for j := 0; j < len(nums); j++ {
			temp = nums[j]
			nums[j] = previous
			previous = temp
		}
	}
}

func rotate2(nums []int, k int) {
	a := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		a[(i+k)%len(nums)] = nums[i]
	}
	for i := 0; i < len(nums); i++ {
		nums[i] = a[i]
	}
	fmt.Println(nums)
}

func rotate3(nums []int, k int) {
	k = k % len(nums)
	copy(nums, append(nums[len(nums)-k:], nums[:len(nums)-k]...))
}

func rotate4(nums []int, k int) {
	k = k % len(nums)
	count := 0
	for start := 0; start < len(nums); start++ {
		current := start
		prev := nums[start]
		for start != current {
			next := (current + k) % len(nums)
			temp := nums[next]
			nums[next] = prev
			prev = temp
			current = next
			count++
		}
	}
}

func rotate5(nums []int, k int) {
	reverse(nums)
	reverse(nums[:k%len(nums)])
	reverse(nums[k%len(nums):])
}

func reverse(arr []int) {
	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-i-1] = arr[len(arr)-i-1], arr[i]
	}
}
