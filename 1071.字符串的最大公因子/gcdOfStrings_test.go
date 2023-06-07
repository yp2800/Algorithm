package leetcode

import "testing"

func Test_gcdOfStrings(t *testing.T) {
	type args struct {
		str1 string
		str2 string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"first test", args{
			str1: "ABCABC",
			str2: "ABC",
		}, "ABC"},
		{"second test", args{
			str1: "ABABAB",
			str2: "ABAB",
		}, "AB"},
		{"third test", args{
			str1: "LEET",
			str2: "CODE",
		}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := gcdOfStrings(tt.args.str1, tt.args.str2); got != tt.want {
				t.Errorf("gcdOfStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
