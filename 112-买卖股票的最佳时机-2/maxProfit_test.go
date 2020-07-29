package leetcode

import "testing"

func Test_maxProfit(t *testing.T) {
	type args struct {
		prices []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"first test", args{
			prices: []int{7, 1, 5, 3, 6, 4},
		}, 7},
		{"second test", args{
			prices: []int{1, 2, 3, 4, 5},
		}, 4},
		{"third test", args{
			prices: []int{7, 6, 4, 3, 1},
		}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit(tt.args.prices); got != tt.want {
				t.Errorf("maxProfit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calculate2(t *testing.T) {
	type args struct {
		prices []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"first test", args{
			prices: []int{7, 1, 5, 3, 6, 4},
		}, 7},
		{"second test", args{
			prices: []int{1, 2, 3, 4, 5},
		}, 4},
		{"third test", args{
			prices: []int{7, 6, 4, 3, 1},
		}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculate2(tt.args.prices); got != tt.want {
				t.Errorf("calculate2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calculate3(t *testing.T) {
	type args struct {
		prices []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"first test", args{
			prices: []int{7, 1, 5, 3, 6, 4},
		}, 7},
		{"second test", args{
			prices: []int{1, 2, 3, 4, 5},
		}, 4},
		{"third test", args{
			prices: []int{7, 6, 4, 3, 1},
		}, 0},
		{"fourth test", args{
			prices: []int{1, 7, 2, 3, 6, 7, 6, 6},
		}, 11},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculate3(tt.args.prices); got != tt.want {
				t.Errorf("calculate3() = %v, want %v", got, tt.want)
			}
		})
	}
}
