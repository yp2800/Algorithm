package leetcode

import "testing"

func Test_merge(t *testing.T) {
	type args struct {
		nums1 []int
		m     int
		nums2 []int
		n     int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{"first test", args{
			nums1: []int{1, 2, 3},
			m:     3,
			nums2: []int{2, 5, 6},
			n:     3,
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}
