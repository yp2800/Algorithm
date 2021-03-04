package leetcode

import "testing"

func Test_isAnagram(t *testing.T) {
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
			s: "anagram",
			t: "nagaram",
		}, true},
		{"second test", args{
			s: "rat",
			t: "car",
		}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isAnagram(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("isAnagram() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isAnagram2(t *testing.T) {
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
			s: "anagram",
			t: "nagaram",
		}, true},
		{"second test", args{
			s: "rat",
			t: "car",
		}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isAnagram2(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("isAnagram2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isAnagram3(t *testing.T) {
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
			s: "anagram",
			t: "nagaram",
		}, true},
		{"second test", args{
			s: "rat",
			t: "car",
		}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isAnagram3(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("isAnagram3() = %v, want %v", got, tt.want)
			}
		})
	}
}
