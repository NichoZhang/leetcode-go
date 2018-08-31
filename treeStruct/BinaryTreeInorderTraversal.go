package main

/**
* https://leetcode.com/problems/binary-tree-inorder-traversal/description/
*
* Input: [1,null,2,3]
*    1
*     \
*      2
*     /
*    3
*
* Output: [1,3,2]
*
* 68 / 68 test cases passed.
* Runtime:0ms
* beats 100%
* https://leetcode.com/submissions/detail/172885016/
**/

// TreeNode struct
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	// 根节点数据初始化
	// 计数栈初始化
	nodes, stack := []int{}, []*TreeNode{}

	for len(stack) > 0 || root != nil {
		// 如果左子树存在
		if root != nil {
			stack = append(stack, root)
			root = root.Left
		} else {
			root = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			nodes = append(nodes, root.Val)
			root = root.Right
		}
	}

	return nodes
}
