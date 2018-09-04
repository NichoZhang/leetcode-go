package main

import "testing"

func TestBinaryLevelOrderBottomTraversal1(t *testing.T) {
	treeNode := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val: 5,
					},
				},
			},
		},
	}

	t.Log(levelOrderBottom(treeNode))
}
