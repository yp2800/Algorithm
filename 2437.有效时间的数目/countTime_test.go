package leetcode

import "testing"

func Test_countTime(t *testing.T) {
	type args struct {
		time string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"first test", args{time: "?5:00"}, 2},
		{"second test", args{time: "0?:0?"}, 100},
		{"third test", args{time: "??:??"}, 1440},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countTime(tt.args.time); got != tt.want {
				t.Errorf("countTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_countTime2(t *testing.T) {
	type args struct {
		time string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"first test", args{time: "?5:00"}, 2},
		{"second test", args{time: "0?:0?"}, 100},
		{"third test", args{time: "??:??"}, 1440},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countTime2(tt.args.time); got != tt.want {
				t.Errorf("countTime2() = %v, want %v", got, tt.want)
			}
		})
	}
}
