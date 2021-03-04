package leetcode

import (
	"reflect"
	"testing"
)

var (
	root = &TreeNode{
		Val: 6,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   0,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val:   3,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val:   5,
					Left:  nil,
					Right: nil,
				},
			},
		},
		Right: &TreeNode{
			Val: 8,
			Left: &TreeNode{
				Val:   7,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   9,
				Left:  nil,
				Right: nil,
			},
		},
	}
)

func Test_lowestCommonAncestor(t *testing.T) {
	type args struct {
		root *TreeNode
		p    *TreeNode
		q    *TreeNode
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		// TODO: Add test cases.
		{"first test", args{
			root: root,
			p:    root.Left,
			q:    root.Right,
		}, root},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lowestCommonAncestor(tt.args.root, tt.args.p, tt.args.q); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("lowestCommonAncestor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_lowestCommonAncestor2(t *testing.T) {
	type args struct {
		root *TreeNode
		p    *TreeNode
		q    *TreeNode
	}
	tests := []struct {
		name         string
		args         args
		wantAncestor *TreeNode
	}{
		// TODO: Add test cases.
		{"first test", args{
			root: root,
			p:    root.Left,
			q:    root.Right,
		}, root},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAncestor := lowestCommonAncestor2(tt.args.root, tt.args.p, tt.args.q); !reflect.DeepEqual(gotAncestor, tt.wantAncestor) {
				t.Errorf("lowestCommonAncestor2() = %v, want %v", gotAncestor, tt.wantAncestor)
			}
		})
	}
}