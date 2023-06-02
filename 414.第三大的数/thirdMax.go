package leetcode

import (
	//"github.com/emirpasic/gods/trees/redblacktree"
	"sort"
)

func thirdMax(nums []int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	for i, diff := 1, 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			diff++
			if diff == 3 {
				return nums[i]
			}
		}
	}
	return nums[0]
}

//func thirdMax2(nums []int) int {
//	t := redblacktree.NewWithIntComparator()
//	for _, num := range nums {
//		t.Put(num, nil)
//		if t.Size() > 3 {
//			t.Remove(t.Left().Key)
//		}
//	}
//	if t.Size() == 3 {
//		return t.Left().Key.(int)
//	}
//	return t.Right().Key.(int)
//}

func thirdMax3(nums []int) int {
	var a, b, c *int
	for _, num := range nums {
		num := num
		if a == nil || num > *a {
			a, b, c = &num, a, b
		} else if *a > num && (b == nil || num > *b) {
			b, c = &num, b
		} else if b != nil && *b > num && (c == nil || num > *c) {
			c = &num
		}
	}
	if c == nil {
		return *a
	}
	return *c
}
