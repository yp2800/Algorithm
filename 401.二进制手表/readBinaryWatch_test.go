package leetcode

import (
	"reflect"
	"testing"
)

func Test_readBinaryWatch(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
		{"first test", args{num: 1}, []string{"1:00", "2:00", "4:00", "8:00", "0:01", "0:02", "0:04", "0:08", "0:16", "0:32"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := readBinaryWatch(tt.args.num); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("readBinaryWatch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_readBinaryWatch2(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
		{"first test", args{num: 1}, []string{"1:00", "2:00", "4:00", "8:00", "0:01", "0:02", "0:04", "0:08", "0:16", "0:32"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := readBinaryWatch2(tt.args.num); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("readBinaryWatch2() = %v, want %v", got, tt.want)
			}
		})
	}
}
