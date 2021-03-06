package leetcode

import "testing"

func Test_isSubsequence(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"first test", args{
			s: "abc",
			t: "ahbgdc",
		}, true},
		{"second test", args{
			s: "axc",
			t: "ahbgdc",
		}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSubsequence(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("isSubsequence() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isSubsequence2(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"first test", args{
			s: "abc",
			t: "ahbgdc",
		}, true},
		{"second test", args{
			s: "axc",
			t: "ahbgdc",
		}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSubsequence2(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("isSubsequence2() = %v, want %v", got, tt.want)
			}
		})
	}
}
