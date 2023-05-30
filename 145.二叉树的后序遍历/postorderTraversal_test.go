package leetcode

import (
	"reflect"
	"testing"
)

func Test_postorderTraversal(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{"first test", args{root: &TreeNode{
			Val:  1,
			Left: nil,
			Right: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val:   3,
					Left:  nil,
					Right: nil,
				},
				Right: nil,
			},
		}}, []int{3, 2, 1}},
		{"second test", args{root: nil}, []int{}},
		{"third test", args{root: &TreeNode{Val: 1}}, []int{1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := postorderTraversal(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("postorderTraversal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_postorderTraversal2(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{"first test", args{root: &TreeNode{
			Val:  1,
			Left: nil,
			Right: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val:   3,
					Left:  nil,
					Right: nil,
				},
				Right: nil,
			},
		}}, []int{3, 2, 1}},
		{"second test", args{root: nil}, []int{}},
		{"third test", args{root: &TreeNode{Val: 1}}, []int{1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := postorderTraversal2(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("postorderTraversal2() = %v, want %v", got, tt.want)
			}
		})
	}
}
