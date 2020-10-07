package leetcode

import "testing"

func Test_reverseBits(t *testing.T) {
	type args struct {
		num uint32
	}
	tests := []struct {
		name string
		args args
		want uint32
	}{
		// TODO: Add test cases.
		{"first test", args{
			num: 43261596,
		}, 964176192},
		{"second test", args{
			num: 4294967293,
		}, 3221225471},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseBits(tt.args.num); got != tt.want {
				t.Errorf("reverseBits() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reverseBits2(t *testing.T) {
	type args struct {
		num uint32
	}
	tests := []struct {
		name string
		args args
		want uint32
	}{
		// TODO: Add test cases.
		{"first test", args{
			num: 43261596,
		}, 964176192},
		{"second test", args{
			num: 4294967293,
		}, 3221225471},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseBits2(tt.args.num); got != tt.want {
				t.Errorf("reverseBits2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reverseBits3(t *testing.T) {
	type args struct {
		num uint32
	}
	tests := []struct {
		name string
		args args
		want uint32
	}{
		// TODO: Add test cases.
		{"first test", args{
			num: 43261596,
		}, 964176192},
		{"second test", args{
			num: 4294967293,
		}, 3221225471},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseBits3(tt.args.num); got != tt.want {
				t.Errorf("reverseBits3() = %v, want %v", got, tt.want)
			}
		})
	}
}
