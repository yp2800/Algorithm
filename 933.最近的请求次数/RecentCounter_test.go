package leetcode_933

import (
	"testing"
)

func TestRecentCounter_Ping(t *testing.T) {
	type args struct {
		t int
	}
	tests := []struct {
		name string
		this RecentCounter
		args args
		want int
	}{
		// TODO: Add test cases.
		{"first test", RecentCounter{1, 100, 3001}, args{3002}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.this.Ping(tt.args.t); got != tt.want {
				t.Errorf("Ping() = %v, want %v", got, tt.want)
			}
		})
	}
}
