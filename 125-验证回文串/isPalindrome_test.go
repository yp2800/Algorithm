package leetcode

import "testing"

func Test_isPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"first test", args{
			s: "A man, a plan, a canal: Panama",
		}, true},
		{"second test", args{
			s: "race a car",
		}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.s); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isPalindrome2(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"first test", args{
			s: "A man, a plan, a canal: Panama",
		}, true},
		{"second test", args{
			s: "race a car",
		}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome2(tt.args.s); got != tt.want {
				t.Errorf("isPalindrome2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isPalindrome3(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"first test", args{
			s: "A man, a plan, a canal: Panama",
		}, true},
		{"second test", args{
			s: "race a car",
		}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome3(tt.args.s); got != tt.want {
				t.Errorf("isPalindrome3() = %v, want %v", got, tt.want)
			}
		})
	}
}
