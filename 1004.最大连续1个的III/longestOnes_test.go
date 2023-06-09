package leetcode

import "testing"

func Test_longestOnes(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"first  test", args{
			nums: []int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0},
			k:    2,
		}, 6},
		{"second test", args{
			nums: []int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1},
			k:    3,
		}, 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestOnes(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("longestOnes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_longestOnes2(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"first  test", args{
			nums: []int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0},
			k:    2,
		}, 6},
		{"second test", args{
			nums: []int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1},
			k:    3,
		}, 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestOnes2(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("longestOnes2() = %v, want %v", got, tt.want)
			}
		})
	}
}
