package leetcode

import (
	"testing"
)

func Test_reverseWords(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"first  test", args{s: "the sky is blue"}, "blue is sky the"},
		{"second test", args{s: "  hello world  "}, "world hello"},
		{"third test", args{s: "a good   example"}, "example good a"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseWords(tt.args.s); got != tt.want {
				t.Errorf("reverseWords() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reverseWords2(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"first  test", args{s: "the sky is blue"}, "blue is sky the"},
		{"second test", args{s: "  hello world  "}, "world hello"},
		{"third test", args{s: "a good   example"}, "example good a"},
		{"fourth test", args{s: "  Bob    Loves  Alice   "}, "Alice Loves Bob"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseWords2(tt.args.s); got != tt.want {
				t.Errorf("reverseWords2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reverseWords3(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"first  test", args{s: "the sky is blue"}, "blue is sky the"},
		{"second test", args{s: "  hello world  "}, "world hello"},
		{"third test", args{s: "a good   example"}, "example good a"},
		{"fourth test", args{s: "  Bob    Loves  Alice   "}, "Alice Loves Bob"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseWords3(tt.args.s); got != tt.want {
				t.Errorf("reverseWords3() = %v, want %v", got, tt.want)
			}
		})
	}
}
