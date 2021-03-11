package leetcode

import "testing"

func Test_reverseVowels(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"first test", args{s: "hello"}, "holle"},
		{"second test", args{s: "leetcode"}, "leotcede"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseVowels(tt.args.s); got != tt.want {
				t.Errorf("reverseVowels() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reverseVowels2(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"first test", args{s: "hello"}, "holle"},
		{"second test", args{s: "leetcode"}, "leotcede"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseVowels2(tt.args.s); got != tt.want {
				t.Errorf("reverseVowels2() = %v, want %v", got, tt.want)
			}
		})
	}
}
