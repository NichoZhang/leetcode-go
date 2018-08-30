package main

/**
* https://leetcode.com/problems/binary-tree-preorder-traversal/description/
*
* Input: [1,null,2,3]
*    1
*     \
*      2
*     /
*    3
*
* Output: [1,2,3]
*
* 68 / 68 test cases passed.
* Runtime:0ms
* beats 100%
* https://leetcode.com/submissions/detail/172657199/
**/

// TreeNode denfined TreeNode
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	nodes, stack := []int{}, []*TreeNode{root}

	for len(stack) > 0 {

		if root != nil {
			nodes = append(nodes, root.Val)

			if root.Right != nil {
				stack = append(stack, root.Right)
			}

			root = root.Left
		} else {
			root = stack[len(stack)-1]
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		}
	}
	return nodes
}
