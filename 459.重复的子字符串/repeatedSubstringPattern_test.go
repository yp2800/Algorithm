package leetcode

import "testing"

func Test_repeatedSubstringPattern(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"first test", args{s: "abab"}, true},
		{"second test", args{s: "aba"}, false},
		{"first test", args{s: "abcabcabcabc"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := repeatedSubstringPattern(tt.args.s); got != tt.want {
				t.Errorf("repeatedSubstringPattern() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_repeatedSubstringPattern2(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"first test", args{s: "abab"}, true},
		{"second test", args{s: "aba"}, false},
		{"first test", args{s: "abcabcabcabc"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := repeatedSubstringPattern2(tt.args.s); got != tt.want {
				t.Errorf("repeatedSubstringPattern2() = %v, want %v", got, tt.want)
			}
		})
	}
}
