package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// 双指针翻转
func reverseList(head *ListNode) *ListNode {
	cur := head
	var pre *ListNode
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}
