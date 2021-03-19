package leetcode

import "testing"

func Test_canConstruct(t *testing.T) {
	type args struct {
		ransomNote string
		magazine   string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"first test", args{
			ransomNote: "a",
			magazine:   "b",
		}, false},
		{"second test", args{
			ransomNote: "aa",
			magazine:   "ab",
		}, false},
		{"third test", args{
			ransomNote: "aa",
			magazine:   "aab",
		}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canConstruct(tt.args.ransomNote, tt.args.magazine); got != tt.want {
				t.Errorf("canConstruct() = %v, want %v", got, tt.want)
			}
		})
	}
}
