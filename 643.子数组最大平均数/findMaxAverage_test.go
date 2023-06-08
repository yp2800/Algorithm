package leetcode

import "testing"

func Test_findMaxAverage(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
		{"first test", args{
			nums: []int{1, 12, -5, -6, 50, 3},
			k:    4,
		}, 12.75},
		{"second test", args{
			nums: []int{5},
			k:    1,
		}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMaxAverage(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("findMaxAverage() = %v, want %v", got, tt.want)
			}
		})
	}
}
