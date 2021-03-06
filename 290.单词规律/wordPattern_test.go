package leetcode

import "testing"

func Test_wordPattern(t *testing.T) {
	type args struct {
		pattern string
		s       string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"first test", args{
			pattern: "abba",
			s:       "dog cat cat dog",
		}, true},
		{"second test", args{
			pattern: "abba",
			s:       "dog cat cat fish",
		}, false},
		{"third test", args{
			pattern: "aaaa",
			s:       "dog cat cat fish",
		}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := wordPattern(tt.args.pattern, tt.args.s); got != tt.want {
				t.Errorf("wordPattern() = %v, want %v", got, tt.want)
			}
		})
	}
}
