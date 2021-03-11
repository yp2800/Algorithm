package leetcode

import "sort"

func intersect(nums1 []int, nums2 []int) []int {
	if len(nums1) > len(nums2) {
		return intersect(nums2, nums1)
	}
	m := map[int]int{}
	for _, num := range nums1 {
		m[num]++
	}
	intersection := []int{}
	for _, nums := range nums2 {
		if m[nums] > 0 {
			intersection = append(intersection, nums)
			m[nums]--
		}
	}
	return intersection
}

func intersect2(nums1, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	length1, length2 := len(nums1), len(nums2)
	index1, index2 := 0, 0

	intersection := []int{}
	if index1 < length1 && index2 < length2 {
		if nums1[index1] < nums2[index2] {
			index1++
		} else if nums1[index1] > nums2[index2] {
			index2++
		} else {
			intersection = append(intersection, nums1[index1])
			index2++
			index1++
		}
	}
	return intersection
}
