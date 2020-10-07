package leetcode

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
	return root == nil || isBalanced(root.Left) && math.Abs(heigh(root.Left)-heigh(root.Right)) < 2 && isBalanced(root.Right)
}

func heigh(node *TreeNode) float64 {
	if node == nil {
		return 0
	}
	return math.Max(heigh(node.Left), heigh(node.Right)) + 1
}

func isBalanced2(root *TreeNode) bool {
	return find(root) != -1
}

func find(root *TreeNode) float64 {
	if root == nil {
		return 0
	}
	l := find(root.Left)
	if l == -1 {
		return -1
	}
	r := find(root.Right)
	if r == -1 {
		return -1
	}
	if math.Abs(l-r) > 1 {
		return -1
	}
	return math.Max(l, r) + 1
}
