package leetcode

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 迭代实现
func levelOrder(root *TreeNode) [][]int {
	ret := [][]int{}
	if root == nil {
		return ret
	}
	q := []*TreeNode{root}
	for i := 0; len(q) > 0; i++ {
		ret = append(ret, []int{})
		p := []*TreeNode{}
		for j := 0; j < len(q); j++ {
			node := q[j]
			ret[i] = append(ret[i], node.Val)
			if node.Left != nil {
				p = append(p, node.Left)
			}
			if node.Right != nil {
				p = append(p, node.Right)
			}
		}
		q = p
	}
	return ret
}

// 递归实现
func levelOrder2(root *TreeNode) [][]int {
	ret := [][]int{}
	if root == nil {
		return ret
	}
	var dfs func(index int, node *TreeNode)
	dfs = func(index int, node *TreeNode) {
		if len(ret) < index {
			ret = append(ret, []int{})
		}
		ret[index-1] = append(ret[index-1], node.Val)
		if node.Left != nil {
			dfs(index+1, node.Left)
		}
		if node.Right != nil {
			dfs(index+1, node.Right)
		}
	}
	dfs(1, root)
	return ret
}
