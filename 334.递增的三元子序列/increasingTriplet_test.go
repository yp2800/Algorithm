package leetcode

import "testing"

func Test_increasingTriplet(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"first test", args{nums: []int{1, 2, 3, 4, 5}}, true},
		{"second test", args{nums: []int{5, 4, 3, 2, 1}}, false},
		{"third test", args{nums: []int{2, 1, 5, 0, 4, 6}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := increasingTriplet(tt.args.nums); got != tt.want {
				t.Errorf("increasingTriplet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_increasingTriplet2(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"first test", args{nums: []int{1, 2, 3, 4, 5}}, true},
		{"second test", args{nums: []int{5, 4, 3, 2, 1}}, false},
		{"third test", args{nums: []int{2, 1, 5, 0, 4, 6}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := increasingTriplet2(tt.args.nums); got != tt.want {
				t.Errorf("increasingTriplet2() = %v, want %v", got, tt.want)
			}
		})
	}
}
