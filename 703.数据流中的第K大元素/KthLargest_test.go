package leetcode_703

import (
	"sort"
	"testing"
)

func TestKthLargest_Add(t *testing.T) {
	type fields struct {
		k        int
		IntSlice sort.IntSlice
	}
	type args struct {
		val int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		// TODO: Add test cases.
		{"first test", fields{
			k:        3,
			IntSlice: sort.IntSlice{4, 5, 8, 2},
		}, args{val: 3}, 4},
		{"second test", fields{
			k:        3,
			IntSlice: sort.IntSlice{4, 5, 8, 2, 3},
		}, args{val: 5}, 5},
		{"third test", fields{
			k:        3,
			IntSlice: sort.IntSlice{4, 5, 8, 2, 3, 5},
		}, args{val: 10}, 5},
		{"fourth test", fields{
			k:        3,
			IntSlice: sort.IntSlice{4, 5, 8, 2, 3, 5, 10},
		}, args{val: 9}, 8},
		{"fifth test", fields{
			k:        3,
			IntSlice: sort.IntSlice{4, 5, 8, 2, 3, 5, 10, 9},
		}, args{val: 4}, 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := Constructor(tt.fields.k, tt.fields.IntSlice)
			kl := &k
			if got := kl.Add(tt.args.val); got != tt.want {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}
