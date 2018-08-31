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

	// 根节点数据初始化
	// 计数栈初始化
	nodes, stack := []int{}, []*TreeNode{root}

	for len(stack) > 0 {

		// 如果根节点存在
		if root != nil {
			// 记录根节点数据
			nodes = append(nodes, root.Val)

			// 如果右子树存在，将右子树压入栈
			if root.Right != nil {
				stack = append(stack, root.Right)
			}

			// 将根替换位左子树
			root = root.Left
		} else { // 如果根节点不存在,从栈顶取出最后压入的树进行遍历
			root = stack[len(stack)-1]
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		}
	}
	return nodes
}
