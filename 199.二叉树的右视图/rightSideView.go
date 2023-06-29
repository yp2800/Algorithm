package leetcode

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 深度优先搜索
func rightSideView(root *TreeNode) []int {
	var ans []int

	var dfs func(node *TreeNode, depth int)
	dfs = func(node *TreeNode, depth int) {
		if node == nil {
			return
		}
		if depth == len(ans) {
			ans = append(ans, node.Val)
		}
		depth++
		dfs(node.Right, depth)
		dfs(node.Left, depth)
	}
	dfs(root, 0)
	return ans
}

// 广度优先搜索
func rightSideView2(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var ans []int
	queue := []*TreeNode{root}
	for len(queue) != 0 {
		n := len(queue)

		for _, node := range queue {
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		ans = append(ans, queue[n-1].Val)
		queue = queue[n:]
	}
	return ans
}
