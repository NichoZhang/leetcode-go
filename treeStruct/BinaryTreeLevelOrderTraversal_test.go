package main

import "testing"

func TestBinaryLevelOrderTraversal1(t *testing.T) {
	treeNode := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val: 15,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}

	t.Log(levelOrder(treeNode))
}
