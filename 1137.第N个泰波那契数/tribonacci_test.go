package leetcode

import "testing"

func Test_tribonacci(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"first test", args{n: 4}, 4},
		{"first test", args{n: 25}, 1389537},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tribonacci(tt.args.n); got != tt.want {
				t.Errorf("tribonacci() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_tribonacci2(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"first test", args{n: 4}, 4},
		{"first test", args{n: 25}, 1389537},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tribonacci2(tt.args.n); got != tt.want {
				t.Errorf("tribonacci2() = %v, want %v", got, tt.want)
			}
		})
	}
}
