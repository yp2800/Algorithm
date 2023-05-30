package leetcode

import (
	"reflect"
	"testing"
)

func Test_preorderTraversal(t *testing.T) {
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
		}}, []int{1, 2, 3}},
		{"second test", args{root: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val:   2,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		}}, []int{1, 2}},
		{"third test", args{root: &TreeNode{
			Val:  1,
			Left: nil,
			Right: &TreeNode{
				Val:   2,
				Left:  nil,
				Right: nil,
			},
		}}, []int{1, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := preorderTraversal(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("preorderTraversal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_preorderTraversal2(t *testing.T) {
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
		}}, []int{1, 2, 3}},
		{"second test", args{root: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val:   2,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		}}, []int{1, 2}},
		{"third test", args{root: &TreeNode{
			Val:  1,
			Left: nil,
			Right: &TreeNode{
				Val:   2,
				Left:  nil,
				Right: nil,
			},
		}}, []int{1, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := preorderTraversal2(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("preorderTraversal2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_preorderTraversal3(t *testing.T) {
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
		}}, []int{1, 2, 3}},
		{"second test", args{root: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val:   2,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		}}, []int{1, 2}},
		{"third test", args{root: &TreeNode{
			Val:  1,
			Left: nil,
			Right: &TreeNode{
				Val:   2,
				Left:  nil,
				Right: nil,
			},
		}}, []int{1, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := preorderTraversal3(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("preorderTraversal3() = %v, want %v", got, tt.want)
			}
		})
	}
}
