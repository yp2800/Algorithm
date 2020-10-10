package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

// 通过哨兵节点去解决它，哨兵节点广泛应用于树和链表中，如伪头、伪尾、标记等，它们是纯功能的，通常不保存任何数据，其主要目的是使链表标准化，如使链表永不为空、永不无头、简化插入和删除。
func removeElements(head *ListNode, val int) *ListNode {
	sentinel := &ListNode{}
	sentinel.Next = head
	prev, curr := sentinel, head
	for curr != nil {
		if curr.Val == val {
			prev.Next = curr.Next
		} else {
			prev = curr
		}
		curr = curr.Next
	}
	return sentinel.Next
}
