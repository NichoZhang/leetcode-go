package main

/**
* https://leetcode.com/problems/binary-tree-postorder-traversal/description/
*
* Example:
* Input: [1,null,2,3]
*    1
*     \
*      2
*     /
*    3
*
* Output: [3,2,1]
*
* 68 / 68 test cases passed.
* Runtime:0ms
* beats 100%
* https://leetcode.com/submissions/detail/173727676/
**/

// TreeNode struct
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postOrder(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	nodes, stack := []int{}, []*TreeNode{}
	visited := &TreeNode{}

	for len(stack) > 0 || root != nil {
		if root != nil {
			stack = append(stack, root)
			root = root.Left
		} else {
			root = stack[len(stack)-1]
			if root.Right != nil && root.Right != visited {
				root = root.Right
			} else {
				nodes = append(nodes, root.Val)
				visited = root
				stack = stack[:len(stack)-1]
				root = nil
			}
		}
	}

	return nodes
}
