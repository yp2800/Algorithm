package leetcode

import (
	"container/heap"
	"sort"
)

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// 暴力求解
func mergeKLists(lists []*ListNode) *ListNode {
	nodes := make([]*ListNode, 0)
	dummyHead := &ListNode{Val: -1}
	point := dummyHead
	for _, list := range lists {
		for list != nil {
			nodes = append(nodes, list)
			list = list.Next
		}
	}

	sort.SliceStable(nodes, func(i, j int) bool {
		return nodes[i].Val < nodes[j].Val
	})
	for i := range nodes {
		point.Next = nodes[i]
		point = point.Next
	}
	return dummyHead.Next
}

// 逐一比较
func mergeKLists2(lists []*ListNode) *ListNode {
	dummyHead := &ListNode{Val: -1}
	tail := dummyHead
	n := len(lists)
	for {
		var minNode *ListNode
		minPointer := -1
		for i := 0; i < n; i++ {
			if lists[i] == nil {
				continue
			}
			if minNode == nil || lists[i].Val < minNode.Val {
				minNode = lists[i]
				minPointer = i
			}
		}
		if minPointer == -1 {
			break
		}
		tail.Next = minNode
		tail = tail.Next
		lists[minPointer] = lists[minPointer].Next

	}
	return dummyHead.Next
}

// 用队列方法优化方法2
func mergeKLists3(lists []*ListNode) *ListNode {
	n := len(lists)
	if n == 0 {
		return nil
	}
	h := IntHeap{}
	heap.Init(&h)
	for _, list := range lists {
		if list != nil {
			heap.Push(&h, list)
		}
	}

	dummyHead := &ListNode{}
	point := dummyHead
	for h.Len() > 0 {
		minNode := heap.Pop(&h).(*ListNode)
		point.Next = minNode
		point = point.Next
		if minNode.Next != nil {
			heap.Push(&h, minNode.Next)
		}
	}
	return dummyHead.Next
}

type IntHeap []*ListNode

func (h IntHeap) Less(i, j int) bool  { return h[i].Val < h[j].Val }
func (h IntHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h IntHeap) Len() int            { return len(h) }
func (h *IntHeap) Push(x interface{}) { *h = append(*h, x.(*ListNode)) }
func (h *IntHeap) Pop() interface{} {
	n := len(*h)
	ans := (*h)[n-1]
	*h = (*h)[:n-1]
	return ans
}

// 逐一合并两条链表
func mergeKLists4(lists []*ListNode) *ListNode {
	var res *ListNode
	for _, list := range lists {
		res = merge2Lists(res, list)
	}
	return res
}

func merge2Lists(l1, l2 *ListNode) *ListNode {
	//合并两条有序链表 — 递归
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		l1.Next = merge2Lists(l1.Next, l2)
		return l1
	}
	l2.Next = merge2Lists(l1, l2.Next)
	return l2
}

func merge2Lists2(l1, l2 *ListNode) *ListNode {
	//合并两条有序链表 — 迭代
	dummyHead := &ListNode{}
	tail := dummyHead
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			tail.Next = l1
			l1 = l1.Next
		} else {
			tail.Next = l2
			l2 = l2.Next
		}
		tail = tail.Next
	}
	if l1 == nil {
		tail.Next = l2
	} else {
		tail.Next = l1
	}
	return dummyHead.Next
}

// 分治合并 (有问题，超出时间限制)
func mergeKLists5(lists []*ListNode) *ListNode {
	n := len(lists)
	return merge(lists, 0, n-1)
}

func merge(lists []*ListNode, l int, r int) *ListNode {
	if l == r {
		return lists[1]
	}
	if l > r {
		return nil
	}
	mid := (l + r) >> 1
	return mergeTwoLists(merge(lists, l, mid), merge(lists, mid+1, r))
}

func mergeTwoLists(a *ListNode, b *ListNode) *ListNode {
	if a == nil || b == nil {
		if a != nil {
			return a
		} else {
			return b
		}
	}
	dummyHead := &ListNode{}
	tail := dummyHead
	aPtr := a
	bPtr := b

	for aPtr != nil && bPtr != nil {
		if aPtr.Val < bPtr.Val {
			tail.Next = aPtr
			aPtr = aPtr.Next
		} else {
			tail.Next = bPtr
			bPtr = bPtr.Next
		}
		tail = tail.Next
	}
	if aPtr != nil {
		tail.Next = aPtr
	} else {
		tail.Next = bPtr
	}
	return dummyHead.Next
}
