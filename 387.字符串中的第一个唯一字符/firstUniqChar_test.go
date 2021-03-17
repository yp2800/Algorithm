package leetcode

import "testing"

func Test_firstUniqChar(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"first test", args{
			s: "leetcode",
		}, 0},
		{"second test", args{s: "loveleetcode"}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := firstUniqChar(tt.args.s); got != tt.want {
				t.Errorf("firstUniqChar() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_firstUniqChar2(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"first test", args{
			s: "leetcode",
		}, 0},
		{"second test", args{s: "loveleetcode"}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := firstUniqChar2(tt.args.s); got != tt.want {
				t.Errorf("firstUniqChar2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_firstUniqChar3(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"first test", args{
			s: "leetcode",
		}, 0},
		{"second test", args{s: "loveleetcode"}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := firstUniqChar3(tt.args.s); got != tt.want {
				t.Errorf("firstUniqChar3() = %v, want %v", got, tt.want)
			}
		})
	}
}
