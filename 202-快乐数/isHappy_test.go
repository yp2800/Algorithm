package leetcode

import "testing"

func Test_isHappy(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"first test", args{
			n: 19,
		}, true},
		{"second test", args{
			n: 4,
		}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isHappy(tt.args.n); got != tt.want {
				t.Errorf("isHappy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isHappy2(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"first test", args{
			n: 19,
		}, true},
		{"second test", args{
			n: 4,
		}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isHappy2(tt.args.n); got != tt.want {
				t.Errorf("isHappy2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isHappy3(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"first test", args{
			n: 19,
		}, true},
		{"second test", args{
			n: 4,
		}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isHappy3(tt.args.n); got != tt.want {
				t.Errorf("isHappy3() = %v, want %v", got, tt.want)
			}
		})
	}
}
