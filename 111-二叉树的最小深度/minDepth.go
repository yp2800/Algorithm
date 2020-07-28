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
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	count := 0
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		count++
		length := len(queue)
		for i := 0; i < length; i++ {
			curNode := queue[i]
			if curNode.Left == nil && curNode.Right == nil {
				return count
			}
			if curNode.Left != nil {
				queue = append(queue, curNode.Left)
			}
			if curNode.Right != nil {
				queue = append(queue, curNode.Right)
			}
		}
		queue = queue[length:]
	}
	return count
}

// dfs
var min int

func minDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	min = 1<<63 - 1
	dfs(root, 1)
	return min
}
func dfs(root *TreeNode, depth int) {
	if root.Left == nil && root.Right == nil {
		if depth < min {
			min = depth
		}
		return
	}
	if root.Left != nil {
		dfs(root.Left, depth+1)
	}
	if root.Right != nil {
		dfs(root.Right, depth+1)
	}
}

func minDepth3(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left != nil && root.Right == nil {
		return minDepth3(root.Left) + 1
	}
	if root.Left == nil && root.Right != nil {
		return minDepth3(root.Right) + 1
	}
	return minInt(minDepth3(root.Left), minDepth3(root.Right)) + 1
}
func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}
