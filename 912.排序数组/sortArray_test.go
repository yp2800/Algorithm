package leetcode

import (
	"reflect"
	"testing"
)

func Test_sortArray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{"first test", args{nums: []int{5, 2, 3, 1}}, []int{1, 2, 3, 5}},
		{"second test", args{nums: []int{5, 1, 1, 2, 0, 0}}, []int{0, 0, 1, 1, 2, 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortArray(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sortArray2(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{"first test", args{nums: []int{5, 2, 3, 1}}, []int{1, 2, 3, 5}},
		{"second test", args{nums: []int{5, 1, 1, 2, 0, 0}}, []int{0, 0, 1, 1, 2, 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortArray2(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortArray2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sortArray3(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{"first test", args{nums: []int{5, 2, 3, 1}}, []int{1, 2, 3, 5}},
		{"second test", args{nums: []int{5, 1, 1, 2, 0, 0}}, []int{0, 0, 1, 1, 2, 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortArray3(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortArray3() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sortArray4(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{"first test", args{nums: []int{5, 2, 3, 1}}, []int{1, 2, 3, 5}},
		{"second test", args{nums: []int{5, 1, 1, 2, 0, 0}}, []int{0, 0, 1, 1, 2, 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortArray4(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortArray4() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sortArray5(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{"first test", args{nums: []int{5, 2, 3, 1}}, []int{1, 2, 3, 5}},
		{"second test", args{nums: []int{5, 1, 1, 2, 0, 0}}, []int{0, 0, 1, 1, 2, 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortArray5(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortArray5() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sortArray6(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{"first test", args{nums: []int{5, 2, 3, 1}}, []int{1, 2, 3, 5}},
		{"second test", args{nums: []int{5, 1, 1, 2, 0, 0}}, []int{0, 0, 1, 1, 2, 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortArray6(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortArray6() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sortArray7(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{"first test", args{nums: []int{5, 2, 3, 1}}, []int{1, 2, 3, 5}},
		{"second test", args{nums: []int{5, 1, 1, 2, 0, 0}}, []int{0, 0, 1, 1, 2, 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortArray7(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortArray7() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sortArray8(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{"first test", args{nums: []int{5, 2, 3, 1}}, []int{1, 2, 3, 5}},
		{"second test", args{nums: []int{5, 1, 1, 2, 0, 0}}, []int{0, 0, 1, 1, 2, 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortArray8(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortArray8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sortArray9(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{"first test", args{nums: []int{5, 2, 3, 1}}, []int{1, 2, 3, 5}},
		{"second test", args{nums: []int{5, 1, 1, 2, 0, 0}}, []int{0, 0, 1, 1, 2, 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortArray9(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortArray9() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sortArray10(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{"first test", args{nums: []int{5, 2, 3, 1}}, []int{1, 2, 3, 5}},
		{"second test", args{nums: []int{5, 1, 1, 2, 0, 0}}, []int{0, 0, 1, 1, 2, 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortArray10(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortArray10() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sortArray11(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{"first test", args{nums: []int{5, 2, 3, 1}}, []int{1, 2, 3, 5}},
		{"second test", args{nums: []int{5, 1, 1, 2, 0, 0}}, []int{0, 0, 1, 1, 2, 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortArray11(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortArray11() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sortArray12(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{"first test", args{nums: []int{5, 2, 3, 1}}, []int{1, 2, 3, 5}},
		{"second test", args{nums: []int{5, 1, 1, 2, 0, 0}}, []int{0, 0, 1, 1, 2, 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortArray12(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortArray12() = %v, want %v", got, tt.want)
			}
		})
	}
}
