package leetcode

import "testing"

func Test_hasCycle(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"first test", args{
			head: &ListNode{
				Val:  3,
				Next: &ListNode{
					Val:  2,
					Next: &ListNode{
						Val:  0,
						Next: &ListNode{
							Val:  -4,
							Next: nil,
						},
					},
				},
			},
		},true},
	}
	tests[0].args.head.Next.Next.Next.Next = tests[0].args.head.Next.Next
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasCycle(tt.args.head); got != tt.want {
				t.Errorf("hasCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}