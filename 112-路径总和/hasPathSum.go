package leetcode

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
// bfs
func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	queue := []*TreeNode{root}
	queueVal := []int{root.Val}
	for len(queue) > 0 {
		curNode := queue[0]
		curNodeVal := queueVal[0]
		queue = queue[1:]
		queueVal = queueVal[1:]
		if curNode.Left == nil && curNode.Right == nil {
			if curNodeVal == sum {
				return true
			}
			continue
		}
		if curNode.Left != nil {
			queue = append(queue, curNode.Left)
			queueVal = append(queueVal, curNode.Left.Val+curNodeVal)
		}
		if curNode.Right != nil {
			queue = append(queue, curNode.Right)
			queueVal = append(queueVal, curNode.Right.Val+curNodeVal)
		}
	}
	return false
}

func hasPathSum2(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return sum == root.Val
	}
	return hasPathSum2(root.Left, sum-root.Val) || hasPathSum2(root.Right, sum-root.Val)
}
