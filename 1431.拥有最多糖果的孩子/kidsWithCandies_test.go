package leetcode

import (
	"reflect"
	"testing"
)

func Test_kidsWithCandies(t *testing.T) {
	type args struct {
		candies      []int
		extraCandies int
	}
	tests := []struct {
		name string
		args args
		want []bool
	}{
		// TODO: Add test cases.
		{"first test", args{
			candies:      []int{2, 3, 5, 1, 3},
			extraCandies: 3,
		}, []bool{true, true, true, false, true}},
		{"second test", args{
			candies:      []int{4, 2, 1, 1, 2},
			extraCandies: 1,
		}, []bool{true, false, false, false, false}},
		{"third test", args{
			candies:      []int{12, 1, 12},
			extraCandies: 10,
		}, []bool{true, false, true}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kidsWithCandies(tt.args.candies, tt.args.extraCandies); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("kidsWithCandies() = %v, want %v", got, tt.want)
			}
		})
	}
}
