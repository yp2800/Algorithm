package leetcode

import "testing"

func Test_pivotIndex(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"first test", args{nums: []int{1, 7, 3, 6, 5, 6}}, 3},
		{"second test", args{nums: []int{1, 2, 3}}, -1},
		{"third test", args{nums: []int{2, 1, -1}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pivotIndex(tt.args.nums); got != tt.want {
				t.Errorf("pivotIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}
