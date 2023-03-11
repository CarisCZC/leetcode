package main

// 两两交换节点

//  Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// 构造虚拟头结点，很实用
func swapPairs(head *ListNode) *ListNode {
	newhead := &ListNode{0, head}
	pre := newhead
	cur := head
	for cur != nil && cur.Next != nil {
		next := cur.Next
		pre.Next = next
		cur.Next = next.Next
		next.Next = cur
		pre = cur
		cur = cur.Next
	}
	return newhead.Next
}
