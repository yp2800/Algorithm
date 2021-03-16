package leetcode

import "testing"

func Test_findTheDifference(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want byte
	}{
		// TODO: Add test cases.
		{"first test", args{
			s: "abcd",
			t: "abcde",
		}, 'e'},
		{"second test", args{
			s: "",
			t: "y",
		}, 'y'},
		{"third test", args{
			s: "a",
			t: "aa",
		}, 'a'},
		{"fourth test", args{
			s: "ae",
			t: "aea",
		}, 'a'},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findTheDifference(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("findTheDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findTheDifference2(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want byte
	}{
		// TODO: Add test cases.
		{"first test", args{
			s: "abcd",
			t: "abcde",
		}, 'e'},
		{"second test", args{
			s: "",
			t: "y",
		}, 'y'},
		{"third test", args{
			s: "a",
			t: "aa",
		}, 'a'},
		{"fourth test", args{
			s: "ae",
			t: "aea",
		}, 'a'},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findTheDifference2(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("findTheDifference2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findTheDifference3(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want byte
	}{
		// TODO: Add test cases.
		{"first test", args{
			s: "abcd",
			t: "abcde",
		}, 'e'},
		{"second test", args{
			s: "",
			t: "y",
		}, 'y'},
		{"third test", args{
			s: "a",
			t: "aa",
		}, 'a'},
		{"fourth test", args{
			s: "ae",
			t: "aea",
		}, 'a'},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findTheDifference3(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("findTheDifference3() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findTheDifference4(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name     string
		args     args
		wantDiff byte
	}{
		// TODO: Add test cases.
		{"first test", args{
			s: "abcd",
			t: "abcde",
		}, 'e'},
		{"second test", args{
			s: "",
			t: "y",
		}, 'y'},
		{"third test", args{
			s: "a",
			t: "aa",
		}, 'a'},
		{"fourth test", args{
			s: "ae",
			t: "aea",
		}, 'a'},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotDiff := findTheDifference4(tt.args.s, tt.args.t); gotDiff != tt.wantDiff {
				t.Errorf("findTheDifference4() = %v, want %v", gotDiff, tt.wantDiff)
			}
		})
	}
}
