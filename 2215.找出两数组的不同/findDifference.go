package leetcode

func findDifference(nums1 []int, nums2 []int) [][]int {
	m1 := make(map[int]struct{})
	m2 := make(map[int]struct{})

	for _, v := range nums1 {
		if _, ok := m1[v]; !ok {
			m1[v] = struct{}{}
		}
	}
	for _, v := range nums2 {
		if _, ok := m2[v]; !ok {
			m2[v] = struct{}{}
		}
	}
	answer := make([][]int, 2)
	for k := range m1 {
		if _, ok := m2[k]; !ok {
			answer[0] = append(answer[0], k)
		}
	}
	for k := range m2 {
		if _, ok := m1[k]; !ok {
			answer[1] = append(answer[1], k)
		}
	}
	return answer
}
