package leetcode

import "testing"

func Test_isIsomorphic(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"first test", args{
			s: "egg",
			t: "add",
		}, true},
		{"second test", args{
			s: "foo",
			t: "bar",
		}, false},
		{"third test", args{
			s: "paper",
			t: "title",
		}, true},
		{"fourth test", args{
			s: "ab",
			t: "aa",
		}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isIsomorphic(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("isIsomorphic() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isIsomorphic2(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"first test", args{
			s: "egg",
			t: "add",
		}, true},
		{"second test", args{
			s: "foo",
			t: "bar",
		}, false},
		{"third test", args{
			s: "paper",
			t: "title",
		}, true},
		{"fourth test", args{
			s: "ab",
			t: "aa",
		}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isIsomorphic2(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("isIsomorphic2() = %v, want %v", got, tt.want)
			}
		})
	}
}
