package leetcode

import (
	"reflect"
	"testing"
)

func Test_plusOne(t *testing.T) {
	type args struct {
		digits []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{"first test", args{
			[]int{1, 2, 3},
		}, []int{1, 2, 4}},
		{"first test", args{
			[]int{4, 3, 2, 1},
		}, []int{4, 3, 2, 2}},
		{"third test", args{
			[]int{9},
		}, []int{1, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := plusOne(tt.args.digits); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("plusOne() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_plusOne2(t *testing.T) {
	type args struct {
		digits []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{"first test", args{
			[]int{1, 2, 3},
		}, []int{1, 2, 4}},
		{"first test", args{
			[]int{4, 3, 2, 1},
		}, []int{4, 3, 2, 2}},
		{"third test", args{
			[]int{9},
		}, []int{1, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := plusOne2(tt.args.digits); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("plusOne2() = %v, want %v", got, tt.want)
			}
		})
	}
}
