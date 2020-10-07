package leetcode

import (
	"reflect"
	"testing"
)

func Test_isPalindrome(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{name: "first test", args: args{
			121,
		}, want: true},
		{name: "second test", args: args{
			-121,
		}, want: false},
		{name: "third test", args: args{
			1212,
		}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.x); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}