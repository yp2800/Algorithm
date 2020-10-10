package leetcode

import (
	"log"
	"testing"
)

// 指针返回值，做值对比，遍历
func Test_removeElements(t *testing.T) {
	type args struct {
		head *ListNode
		val  int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		// TODO: Add test cases.
		{"first test", args{
			head: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 6,
						Next: &ListNode{
							Val: 3,
							Next: &ListNode{
								Val: 4,
								Next: &ListNode{
									Val: 5,
									Next: &ListNode{
										Val:  6,
										Next: nil,
									},
								},
							},
						},
					},
				},
			},
			val: 6,
		}, &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val:  5,
							Next: nil,
						},
					},
				},
			},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := removeElements(tt.args.head, tt.args.val)
			ERR := false
			for got != nil && tt.want != nil {
				if got.Val != tt.want.Val {
					ERR = true
					break
				}
				got = got.Next
				tt.want = tt.want.Next
				if got == nil && tt.want != nil {
					ERR = true
					break
				}
				if got != nil && tt.want == nil {
					ERR = true
					break
				}
			}
			log.Println(ERR)
			if ERR {
				t.Errorf("removeElements() = %v, want %v", got, tt.want)
			}
		})
	}
}
