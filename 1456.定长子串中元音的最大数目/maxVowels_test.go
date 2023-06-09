package leetcode

import "testing"

func Test_maxVowels(t *testing.T) {
	type args struct {
		s string
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"first test", args{
			s: "abciiidef",
			k: 3,
		}, 3},
		{"second test", args{
			s: "aeiou",
			k: 2,
		}, 2},
		{"third test", args{
			s: "leetcode",
			k: 3,
		}, 2},
		{"fourth test", args{
			s: "rhythms",
			k: 4,
		}, 0},
		{"fivth test", args{
			s: "tryhard",
			k: 4,
		}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxVowels(tt.args.s, tt.args.k); got != tt.want {
				t.Errorf("maxVowels() = %v, want %v", got, tt.want)
			}
		})
	}
}
