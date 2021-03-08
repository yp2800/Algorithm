package leetcode

import "testing"

func Test_isPowerOfFour(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"first test", args{n: 16}, true},
		{"second test", args{n: 5}, false},
		{"third test", args{n: 1}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPowerOfFour(tt.args.n); got != tt.want {
				t.Errorf("isPowerOfFour() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isPowerOfFour2(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"first test", args{n: 16}, true},
		{"second test", args{n: 5}, false},
		{"third test", args{n: 1}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPowerOfFour2(tt.args.n); got != tt.want {
				t.Errorf("isPowerOfFour2() = %v, want %v", got, tt.want)
			}
		})
	}
}