package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	var prev, curr *ListNode = nil, slow
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	mid := prev
	for mid != nil {
		if head.Val != mid.Val {
			return false
		}
		mid = mid.Next
		head = head.Next
	}
	return true
}
