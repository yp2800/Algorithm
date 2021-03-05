package leetcode

import (
	"reflect"
	"testing"
)

func Test_binaryTreePaths(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
		{"first test", args{root: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val:  2,
				Left: nil,
				Right: &TreeNode{
					Val:   5,
					Left:  nil,
					Right: nil,
				},
			},
			Right: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
		}}, []string{"1->2->5", "1->3"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := binaryTreePaths(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("binaryTreePaths() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_binaryTreePaths1(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
		{"first test", args{root: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val:  2,
				Left: nil,
				Right: &TreeNode{
					Val:   5,
					Left:  nil,
					Right: nil,
				},
			},
			Right: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
		}}, []string{"1->2->5", "1->3"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := binaryTreePaths2(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("binaryTreePaths2() = %v, want %v", got, tt.want)
			}
		})
	}
}