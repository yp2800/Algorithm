package leetcode

import "testing"

func TestMovingAverage_Next(t *testing.T) {
	type fields struct {
		size int
		sum  int
		nums []int
	}
	type args struct {
		val int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   float64
	}{
		// TODO: Add test cases.
		{"first test", fields{
			size: 3,
			sum:  0,
			nums: []int{},
		}, args{val: 1}, 1},
		{"second test", fields{
			size: 3,
			sum:  1,
			nums: []int{1},
		}, args{val: 10}, 5.5},
		{"third test", fields{
			size: 3,
			sum:  11,
			nums: []int{1, 10},
		}, args{val: 3}, 4.666666666666667},
		{"fourth test", fields{
			size: 3,
			sum:  14,
			nums: []int{1, 10, 3},
		}, args{val: 5}, 6.0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &MovingAverage{
				size: tt.fields.size,
				sum:  tt.fields.sum,
				nums: tt.fields.nums,
			}
			if got := this.Next(tt.args.val); got != tt.want {
				t.Errorf("Next() = %v, want %v", got, tt.want)
			}
		})
	}
}
