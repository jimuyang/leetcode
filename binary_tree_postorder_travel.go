package main

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	var stack list.List
	for node := root; node != nil; node = node.Left {
		stack.PushFront(node)
	}
	res := make([]int, 0)

	// 上次访问的node
	var current *TreeNode
	for stack.Len() > 0 {
		top := stack.Front().Value.(*TreeNode)
		if top.Right == nil || top.Right == current {
			current = top
			res = append(res, current.Val)
			stack.Remove(stack.Front())
		} else if top.Right != nil {
			for node := top.Right; node != nil; node = node.Left {
				stack.PushFront(node)
			}
		}
	}
	return res
}
