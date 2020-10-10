package main

import (
	"reflect"
	"testing"
)

func TestArraySource(t *testing.T) {
	type args struct {
		a []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{"first test", args{a: []int{3, 2, 6, 7, 4}}, []int{3, 2, 6, 7, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cgot := ArraySource(tt.args.a...)
			var got []int
			for i := range cgot {
				got = append(got, i)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ArraySource() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInMemSort(t *testing.T) {
	type args struct {
		in []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{"first test", args{in: []int{3, 2, 6, 7, 4}}, []int{2, 3, 4, 6, 7}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cin := make(chan int)
			go func() {
				for i := range tt.args.in {
					cin <- i
				}
			}()
			cgot := InMemSort(cin)
			var got []int
			for i := range cgot {
				got = append(got, i)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InMemSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
