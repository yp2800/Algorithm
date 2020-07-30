package leetcode

import "testing"

func Test_compareVersion(t *testing.T) {
	type args struct {
		version1 string
		version2 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"first test", args{
			version1: "0.1",
			version2: "1.1",
		}, -1},
		{"second test", args{
			version1: "1.0.1",
			version2: "1",
		}, 1},
		{"third test", args{
			version1: "7.5.2.4",
			version2: "7.5.3",
		}, -1},
		{"fourth test", args{
			version1: "1.01",
			version2: "1.001",
		}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := compareVersion(tt.args.version1, tt.args.version2); got != tt.want {
				t.Errorf("compareVersion() = %v, want %v", got, tt.want)
			}
		})
	}
}
