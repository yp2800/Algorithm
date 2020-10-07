package leetcode

import "testing"

func Test_largestRectangleArea(t *testing.T) {
	type args struct {
		heights []int
	}
	tests := []struct {
		name string
		args args
		 want int
	}{
		// TODO: Add test cases.
		{name: "first test", args: args{
			[]int{1, 1},
		}, want: 2},
		{name: "second test", args: args{
			[]int{1},
		}, want: 1},
		{name: "third test", args: args{
			[]int{2, 1, 5, 6, 2, 3},
		}, want: 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestRectangleArea(tt.args.heights); got != tt.want {
				t.Errorf("largestRectangleArea() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_largestRectangleArea2(t *testing.T) {
	type args struct {
		heights []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{name: "first test", args: args{
			[]int{1, 1},
		}, want: 2},
		{name: "second test", args: args{
			[]int{1},
		}, want: 1},
		{name: "third test", args: args{
			[]int{2, 1, 5, 6, 2, 3},
		}, want: 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestRectangleArea2(tt.args.heights); got != tt.want {
				t.Errorf("largestRectangleArea2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_largestRectangleArea3(t *testing.T) {
	type args struct {
		heights []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{name: "first test", args: args{
			[]int{1, 1},
		}, want: 2},
		{name: "second test", args: args{
			[]int{1},
		}, want: 1},
		{name: "third test", args: args{
			[]int{2, 1, 5, 6, 2, 3},
		}, want: 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestRectangleArea3(tt.args.heights); got != tt.want {
				t.Errorf("largestRectangleArea3() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_largestRectangleArea5(t *testing.T) {
	type args struct {
		heights []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{name: "first test", args: args{
			[]int{1, 1},
		}, want: 2},
		{name: "second test", args: args{
			[]int{1},
		}, want: 1},
		{name: "third test", args: args{
			[]int{2, 1, 5, 6, 2, 3},
		}, want: 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestRectangleArea5(tt.args.heights); got != tt.want {
				t.Errorf("largestRectangleArea5() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_min(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{name: "first test", args: args{
			1,2,
		},want: 1},
		{name: "second test", args: args{
			122222, 23333,
		}, want: 23333},
		{name: "third test", args: args{
			2342341, 22331242,
		}, want: 2342341},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := min(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("min() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_max(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{name: "first test", args: args{
			1,2,
		},want: 2},
		{name: "second test", args: args{
			122222, 23333,
		}, want: 122222},
		{name: "third test", args: args{
			2342341, 22331242,
		}, want: 22331242},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := max(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("max() = %v, want %v", got, tt.want)
			}
		})
	}
}