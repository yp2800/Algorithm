package leetcode

import "testing"

func Test_toHex(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"first test", args{num: 26}, "1a"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toHex(tt.args.num); got != tt.want {
				t.Errorf("toHex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_toHex2(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"first test", args{num: 26}, "1a"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toHex2(tt.args.num); got != tt.want {
				t.Errorf("toHex2() = %v, want %v", got, tt.want)
			}
		})
	}
}
