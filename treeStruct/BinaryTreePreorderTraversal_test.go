package main

import "testing"

func TestPreorderTraversal1(t *testing.T) {
	tn := &TreeNode{
		Val:  1,
		Left: nil,
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
			Right: nil,
		},
	}
	t.Log(preorderTraversal(tn))
}

func TestPreorderTraversal2(t *testing.T) {
	tn := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val: 6,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}
	t.Log(preorderTraversal(tn))
}
