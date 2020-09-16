package leetcode

import (
	"reflect"
	"testing"
)

func Test_ipsort(t *testing.T) {
	type args struct {
		ips []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
		{"first test", args{ips: []string{"1.1.1.1", "1.0.0.1", "2.2.2.2", "2.0.0.1"}}, []string{"2.2.2.2", "2.0.0.1", "1.1.1.1", "1.0.0.1"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ipv4Sort(tt.args.ips); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ipv4Sort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_ipsort2(t *testing.T) {
	type args struct {
		ips []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
		{"first test", args{ips: []string{"1.1.1.1", "1.0.0.1", "2.2.2.2", "2.0.0.1"}}, []string{"2.2.2.2", "2.0.0.1", "1.1.1.1", "1.0.0.1"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ipv4Sort2(tt.args.ips); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ipv4Sort2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_ipv4Sort3(t *testing.T) {
	type args struct {
		ips []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
		{"first test", args{ips: []string{"1.1.1.1", "1.0.0.1", "2.2.2.2", "2.0.0.1"}}, []string{"2.2.2.2", "2.0.0.1", "1.1.1.1", "1.0.0.1"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ipv4Sort3(tt.args.ips); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ipv4Sort3() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_ipv4Sort4(t *testing.T) {
	type args struct {
		ips []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
		{"first test", args{ips: []string{"1.1.1.1", "1.0.0.1", "2.2.2.2", "2.0.0.1"}}, []string{"2.2.2.2", "2.0.0.1", "1.1.1.1", "1.0.0.1"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ipv4Sort4(tt.args.ips); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ipv4Sort4() = %v, want %v", got, tt.want)
			}
		})
	}
}