package leetcode

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归
func preorderTraversal(root *TreeNode) []int {
	var res []int
	var preorder func(node *TreeNode)

	res = []int{}
	preorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		res = append(res, node.Val)
		preorder(node.Left)
		preorder(node.Right)
	}
	preorder(root)
	return res
}

// 迭代
func preorderTraversal2(root *TreeNode) []int {
	var stack []*TreeNode
	var res []int

	stack = []*TreeNode{}
	res = []int{}
	node := root
	for node != nil || len(stack) > 0 {
		for node != nil {
			res = append(res, node.Val)
			stack = append(stack, node)
			node = node.Left
		}
		node = stack[len(stack)-1].Right
		stack = stack[:len(stack)-1]
	}
	return res
}

// Morris 遍历
func preorderTraversal3(root *TreeNode) []int {
	res := []int{}
	var p1, p2 *TreeNode = root, nil
	for p1 != nil {
		p2 = p1.Left
		if p2 != nil {
			for p2.Right != nil && p2.Right != p1 {
				p2 = p2.Right
			}
			if p2.Right == nil {
				res = append(res, p1.Val)
				p2.Right = p1
				p1 = p1.Left
				continue
			}
			p2.Right = nil
		} else {
			res = append(res, p1.Val)
		}
		p1 = p1.Right
	}
	return res
}
