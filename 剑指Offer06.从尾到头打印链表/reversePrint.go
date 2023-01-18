package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func reversePrint(head *ListNode) []int {
	var reverse []int
	for head != nil {
		reverse = append(reverse, head.Val)
		head = head.Next
	}
	for i, j := 0, len(reverse)-1; i < j; {
		reverse[i], reverse[j] = reverse[j], reverse[i]
		i++
		j--
	}
	return reverse
}
