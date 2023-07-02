package leetcode

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// 穿针引线
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	//因为头节点有可能发生变化，使用虚拟头节点可以避免复杂的分类讨论
	dummyNode := &ListNode{Val: -1}
	dummyNode.Next = head

	pre := dummyNode
	//第1步：从虚拟头节点走left-1步，来到left节点的前一个节点
	//建议写在for循环里，语言清晰
	for i := 0; i < left-1; i++ {
		pre = pre.Next
	}

	//第2步：从pre再走right-left+1步，来到right节点
	rightNode := pre
	for i := 0; i < right-left+1; i++ {
		rightNode = rightNode.Next
	}

	//第3步：切断出一个子链表（截取链表）
	leftNode := pre.Next
	curr := rightNode.Next

	//注意：切断链接
	pre.Next = nil
	rightNode.Next = nil

	//第4步：同第206题，反转链表的子区间
	reverseLinkedList(leftNode)

	//第5步：接回到原来的链表中
	pre.Next = rightNode
	leftNode.Next = curr
	return dummyNode.Next
}

func reverseLinkedList(head *ListNode) {
	var pre *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
}

// 一次遍历（穿针引线）反转链表（头插法）
func reverseBetween2(head *ListNode, left, right int) *ListNode {
	dummyNode := &ListNode{Val: -1}
	dummyNode.Next = head
	pre := dummyNode
	for i := 0; i < left-1; i++ {
		pre = pre.Next
	}
	cur := pre.Next
	for i := 0; i < right-left; i++ {
		next := cur.Next
		cur.Next = next.Next
		next.Next = pre.Next
		pre.Next = next
	}
	return dummyNode.Next
}
