package leetcode

import "sort"

func intersection(nums1 []int, nums2 []int) []int {
	m := map[int]struct{}{}
	for _, i := range nums1 {
		if _, ok := m[i]; !ok {
			m[i] = struct{}{}
		}
	}
	m2 := map[int]struct{}{}
	res := []int{}
	for _, i := range nums2 {
		if _, ok := m[i]; ok {
			if _, ok := m2[i]; !ok {
				res = append(res, i)
				m2[i] = struct{}{}
			}
		}
	}
	return res
}

func intersection2(nums1 []int, nums2 []int) (res []int) {
	sort.Ints(nums1)
	sort.Ints(nums2)
	for i, j := 0, 0; i < len(nums1) && j < len(nums2); {
		x, y := nums1[i], nums2[j]
		if x == y {
			if res == nil || x > res[len(res)-1] {
				res = append(res, x)
			}
			i++
			j++
		} else if x < y {
			i++
		} else {
			j++
		}
	}
	return
}
