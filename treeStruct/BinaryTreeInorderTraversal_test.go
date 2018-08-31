package main

import "testing"

func TestInorderTraversal1(t *testing.T) {
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
	t.Log(inorderTraversal(tn))
}

// 1:a
// 2:b
// 3:c
// 4:d
// 5:e
// 6:f
// 7:g
// 8:h
// 9:i
func TestInorderTraversal2(t *testing.T) {
	tn := &TreeNode{
		Val: 6,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 1,
			},
			Right: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 3,
				},
				Right: &TreeNode{
					Val: 5,
				},
			},
		},
		Right: &TreeNode{
			Val:  7,
			Left: nil,
			Right: &TreeNode{
				Val: 9,
				Left: &TreeNode{
					Val: 8,
				},
				Right: nil,
			},
		},
	}
	t.Log(inorderTraversal(tn))
}
