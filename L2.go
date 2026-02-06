package main

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := new(ListNode)
	cur := head

	//它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字
	jin := 0

	for l1 != nil || l2 != nil || jin > 0 {
		s := jin
		if l1 != nil {
			s += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			s += l2.Val
			l2 = l2.Next
		}
		if s >= 10 {
			s = s - 10
			jin = 1
		} else {
			jin = 0
		}
		cur.Next = &ListNode{Val: s}
		cur = cur.Next

	}
	return head.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
