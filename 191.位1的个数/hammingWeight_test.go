package leetcode

import "testing"

func Test_hammingWeight(t *testing.T) {
	type args struct {
		num uint32
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"first test", args{
			num: 11,
		}, 3},
		{"second test", args{
			num: 128,
		}, 1},
		{"third test", args{
			num: 4294967293,
		}, 31},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hammingWeight(tt.args.num); got != tt.want {
				t.Errorf("hammingWeight() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_hammingWeight2(t *testing.T) {
	type args struct {
		num uint32
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"first test", args{
			num: 11,
		}, 3},
		{"second test", args{
			num: 128,
		}, 1},
		{"third test", args{
			num: 4294967293,
		}, 31},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hammingWeight2(tt.args.num); got != tt.want {
				t.Errorf("hammingWeight2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_hammingWeight3(t *testing.T) {
	type args struct {
		num uint32
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"first test", args{
			num: 11,
		}, 3},
		{"second test", args{
			num: 128,
		}, 1},
		{"third test", args{
			num: 4294967293,
		}, 31},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hammingWeight3(tt.args.num); got != tt.want {
				t.Errorf("hammingWeight3() = %v, want %v", got, tt.want)
			}
		})
	}
}
