package leetcode

import (
	"reflect"
	"testing"
)

func TestOrderedStream_Insert(t *testing.T) {
	type fields struct {
		ptr    int
		stream []string
	}
	type args struct {
		idKey int
		value string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []string
	}{
		// TODO: Add test cases.
		{"first test", fields{
			ptr:    1,
			stream: []string{"", "", "", "", "", ""},
		}, args{
			idKey: 3,
			value: "ccccc",
		}, []string{}},
		{"second test", fields{
			ptr:    1,
			stream: []string{"", "", "", "ccccc", "", ""},
		}, args{
			idKey: 1,
			value: "aaaaa",
		}, []string{"aaaaa"}},
		{"third test", fields{
			ptr:    2,
			stream: []string{"", "aaaaa", "", "ccccc", "", ""},
		}, args{
			idKey: 2,
			value: "bbbbb",
		}, []string{"bbbbb", "ccccc"}},
		{"fourth test", fields{
			ptr:    4,
			stream: []string{"", "aaaaa", "bbbbb", "ccccc", "", ""},
		}, args{
			idKey: 5,
			value: "eeeee",
		}, []string{}},
		{"fifth test", fields{
			ptr:    4,
			stream: []string{"", "aaaaa", "bbbbb", "ccccc", "", "eeeee"},
		}, args{
			idKey: 4,
			value: "ddddd",
		}, []string{"ddddd", "eeeee"}},
	}
	//假定 n := 5 + 1,第一个空出来
	n := 6
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := Constructor(n)
			this.ptr = tt.fields.ptr
			this.stream = tt.fields.stream
			if got := this.Insert(tt.args.idKey, tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Insert() = %v, want %v", got, tt.want)
			}
		})
	}
}
