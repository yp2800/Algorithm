package leetcode

import "testing"

func Test_missingNumber(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"first test", args{nums: []int{9, 6, 4, 2, 3, 5, 7, 0, 1}}, 8},
		{"second test", args{nums: []int{3, 0, 1}}, 2},
		{"third test", args{nums: []int{0, 1}}, 2},
		{"fourth test", args{nums: []int{0}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := missingNumber(tt.args.nums); got != tt.want {
				t.Errorf("missingNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_missingNumber2(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"first test", args{nums: []int{9, 6, 4, 2, 3, 5, 7, 0, 1}}, 8},
		{"second test", args{nums: []int{3, 0, 1}}, 2},
		{"third test", args{nums: []int{0, 1}}, 2},
		{"fourth test", args{nums: []int{0}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := missingNumber2(tt.args.nums); got != tt.want {
				t.Errorf("missingNumber2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_missingNumber3(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"first test", args{nums: []int{9, 6, 4, 2, 3, 5, 7, 0, 1}}, 8},
		{"second test", args{nums: []int{3, 0, 1}}, 2},
		{"third test", args{nums: []int{0, 1}}, 2},
		{"fourth test", args{nums: []int{0}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := missingNumber3(tt.args.nums); got != tt.want {
				t.Errorf("missingNumber3() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_missingNumber4(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"first test", args{nums: []int{9, 6, 4, 2, 3, 5, 7, 0, 1}}, 8},
		{"second test", args{nums: []int{3, 0, 1}}, 2},
		{"third test", args{nums: []int{0, 1}}, 2},
		{"fourth test", args{nums: []int{0}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := missingNumber4(tt.args.nums); got != tt.want {
				t.Errorf("missingNumber4() = %v, want %v", got, tt.want)
			}
		})
	}
}
