package main

/**
* https://leetcode.com/problems/binary-tree-level-order-traversal-ii/description/
* For example:
* Given binary tree [3,9,20,null,null,15,7],
*     3
*    / \
*   9  20
*     /  \
*    15   7
* return its bottom-up level order traversal as:
* [
*   [15,7],
*   [9,20],
*   [3]
* ]
*
* 34 / 34 test cases passed.
* Runtime:4ms
* https://leetcode.com/submissions/detail/173477677/
**/

// TreeNode struct
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	ret, stack := [][]int{}, []*TreeNode{root}
	for len(stack) > 0 {
		lenS := len(stack)
		var list []int
		for i := 0; i < lenS; i++ {
			n := stack[i]
			list = append(list, n.Val)
			if n.Left != nil {
				stack = append(stack, n.Left)
			}
			if n.Right != nil {
				stack = append(stack, n.Right)
			}
		}
		stack = stack[lenS:]
		ret = append(ret, list)
	}

	lenR := len(ret)
	for left, right := 0, lenR-1; left < right; left, right = left+1, right-1 {
		ret[left], ret[right] = ret[right], ret[left]
	}
	return ret
}
