package leetcode

import (
	"reflect"
	"testing"
)

func Test_levelOrder(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// TODO: Add test cases.
		{"first test", args{root: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val:   9,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val: 20,
				Left: &TreeNode{
					Val:   15,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val:   7,
					Left:  nil,
					Right: nil,
				},
			},
		}}, [][]int{{3}, {9, 20}, {15, 7}}},
		{"second test", args{root: &TreeNode{
			Val:   1,
			Left:  nil,
			Right: nil,
		}}, [][]int{{1}}},
		{"third test", args{root: nil}, [][]int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := levelOrder(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("levelOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_levelOrder2(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// TODO: Add test cases.
		{"first test", args{root: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val:   9,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val: 20,
				Left: &TreeNode{
					Val:   15,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val:   7,
					Left:  nil,
					Right: nil,
				},
			},
		}}, [][]int{{3}, {9, 20}, {15, 7}}},
		{"second test", args{root: &TreeNode{
			Val:   1,
			Left:  nil,
			Right: nil,
		}}, [][]int{{1}}},
		{"third test", args{root: nil}, [][]int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := levelOrder2(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("levelOrder2() = %v, want %v", got, tt.want)
			}
		})
	}
}
