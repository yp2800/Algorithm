package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	pA := headA
	pB := headB
	flagA := true
	flagB := true
	for {
		if pA == pB {
			return pA
		}

		pA = pA.Next
		pB = pB.Next

		if pA == nil && flagA {
			pA = headB
			flagA = false
		}
		if pB == nil && flagB {
			pB = headA
			flagB = false
		}
	}
	return nil
}
