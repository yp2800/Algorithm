package leetcode

import "testing"

func Test_countPrimes(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"first test", args{10},4},
		{"second test", args{0},0},
		{"third test", args{1},0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countPrimes(tt.args.n); got != tt.want {
				t.Errorf("countPrimes() = %v, want %v", got, tt.want)
			}
		})
	}
}
