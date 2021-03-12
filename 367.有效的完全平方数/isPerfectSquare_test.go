package leetcode

import "testing"

func
Test_isPerfectSquare(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"first test", args{
			num: 16,
		}, true},
		{"second test", args{
			num: 14,
		}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPerfectSquare(tt.args.num); got != tt.want {
				t.Errorf("isPerfectSquare() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isPerfectSquare2(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"first test", args{
			num: 16,
		}, true},
		{"second test", args{
			num: 14,
		}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPerfectSquare2(tt.args.num); got != tt.want {
				t.Errorf("isPerfectSquare2() = %v, want %v", got, tt.want)
			}
		})
	}
}