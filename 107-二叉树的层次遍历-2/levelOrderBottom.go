package leetcode

import (
	"container/list"
)

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
func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var result [][]int
	nodes := []*TreeNode{root}
	for len(nodes) > 0 {
		length := len(nodes)
		list := make([]int, 0)
		for i := 0; i < length; i++ {
			curNode := nodes[i]
			list = append(list, curNode.Val)
			if curNode.Left != nil {
				nodes = append(nodes, curNode.Left)
			}
			if curNode.Right != nil {
				nodes = append(nodes, curNode.Right)
			}
		}
		result = append([][]int{list}, result...)
		nodes = nodes[length:]
	}
	return result
}

func levelOrderBottom2(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var result [][]int
	list := list.New()
	list.PushFront(root)
	for list.Len() > 0 {
		currentLevel := []int{}
		listLength := list.Len()
		for i := 0; i < listLength; i++ {
			node := list.Remove(list.Back()).(*TreeNode)
			currentLevel = append(currentLevel, node.Val)
			if node.Left != nil {
				list.PushFront(node.Left)
			}
			if node.Right != nil {
				list.PushFront(node.Right)
			}
		}
		result = append([][]int{currentLevel}, result...)
	}
	return result
}
