package leetcode

import (
	"reflect"
	"testing"
)

func Test_productExceptSelf(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{"first test", args{nums: []int{1, 2, 3, 4}}, []int{24, 12, 8, 6}},
		{"second test", args{nums: []int{-1, 1, 0, -3, 3}}, []int{0, 0, 9, 0, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := productExceptSelf(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("productExceptSelf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_productExceptSelf2(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{"first test", args{nums: []int{1, 2, 3, 4}}, []int{24, 12, 8, 6}},
		{"second test", args{nums: []int{-1, 1, 0, -3, 3}}, []int{0, 0, 9, 0, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := productExceptSelf2(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("productExceptSelf2() = %v, want %v", got, tt.want)
			}
		})
	}
}
