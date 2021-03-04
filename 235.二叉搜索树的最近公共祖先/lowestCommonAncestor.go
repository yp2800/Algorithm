package leetcode

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) (ancestor *TreeNode) {
	pathP := getPath(root, p)
	pathQ := getPath(root, q)

	for i := 0; i < len(pathP) && i < len(pathQ) && pathP[i] == pathQ[i]; i++ {
		ancestor = pathQ[i]
	}
	return
}

func getPath(root, target *TreeNode) (path []*TreeNode) {
	node := root
	for node != target {
		path = append(path, node)
		if node.Val > target.Val {
			node = node.Left
		} else {
			node = node.Right
		}
	}
	path = append(path, node)
	return
}
func lowestCommonAncestor2(root, p, q *TreeNode) (ancestor *TreeNode) {
	ancestor = root
	for true {
		if p.Val < ancestor.Val && q.Val < ancestor.Val {
			ancestor = ancestor.Left
		} else if p.Val > ancestor.Val && q.Val > ancestor.Val {
			ancestor = ancestor.Right
		} else {
			return
		}
	}
	return
}
