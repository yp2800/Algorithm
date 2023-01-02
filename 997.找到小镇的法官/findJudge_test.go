package leetcode

import "testing"

func Test_findJudge(t *testing.T) {
	type args struct {
		n     int
		trust [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"first test", args{
			n:     2,
			trust: [][]int{{1, 2}},
		}, 2},
		{"second test", args{
			n:     3,
			trust: [][]int{{1, 3}, {2, 3}},
		}, 3},
		{"third test", args{
			n:     3,
			trust: [][]int{{1, 3}, {2, 3}, {3, 1}},
		}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findJudge(tt.args.n, tt.args.trust); got != tt.want {
				t.Errorf("findJudge() = %v, want %v", got, tt.want)
			}
		})
	}
}
