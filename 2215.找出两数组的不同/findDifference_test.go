package leetcode

import (
	"reflect"
	"testing"
)

func Test_findDifference(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// TODO: Add test cases.
		{"first test", args{
			nums1: []int{1, 2, 3},
			nums2: []int{2, 4, 6},
		}, [][]int{{1, 3}, {4, 6}}},
		{"second test", args{
			nums1: []int{1, 2, 3, 3},
			nums2: []int{1, 1, 2, 2},
		}, [][]int{{3}, {}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findDifference(tt.args.nums1, tt.args.nums2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}
