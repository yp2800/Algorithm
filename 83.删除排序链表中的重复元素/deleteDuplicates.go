package leetcode

/**
* Definition for singly-linked list.
* type ListNode struct {
*     Val int
*     Next *ListNode
* }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	point := head.Next
	result := head
	for point != nil {
		if result.Val == point.Val {
			point = point.Next
			result.Next = nil
			continue
		}
		result.Next = point
		result = result.Next
		point = point.Next
	}
	return head
}

func deleteDuplicates2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	head.Next = deleteDuplicates2(head.Next)
	if head.Val == head.Next.Val {
		head.Next = head.Next.Next
	}
	return head
}