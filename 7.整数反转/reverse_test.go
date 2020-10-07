package leetcode

import "testing"

func Test_reverse(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{name: "first test",args: args{
			x: 123,
		}, want: 321},
		{name: "second test",args: args{
			x: 123320,
		}, want: 23321},
		{name: "third test",args: args{
			x: -43221,
		}, want: -12234},
		{name: "fourth test",args: args{
			x: 1534236469,
		}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverse(tt.args.x); got != tt.want {
				t.Errorf("reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}