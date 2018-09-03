package main

// TreeNode struct
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
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
	return ret
}
