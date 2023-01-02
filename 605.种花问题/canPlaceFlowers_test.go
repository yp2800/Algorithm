package leetcode

import "testing"

func Test_canPlaceFlowers(t *testing.T) {
	type args struct {
		flowerbed []int
		n         int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"first test", args{
			flowerbed: []int{1, 0, 0, 0, 0, 1},
			n:         1,
		}, true},
		{"second test", args{
			flowerbed: []int{10, 0, 0, 1},
			n:         2,
		}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canPlaceFlowers(tt.args.flowerbed, tt.args.n); got != tt.want {
				t.Errorf("canPlaceFlowers() = %v, want %v", got, tt.want)
			}
		})
	}
}
