// https://leetcode.cn/problems/reverse-linked-list/
// 翻转链表
package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// 双指针
func reverseList(head *ListNode) *ListNode {
	cur := head
	pre := ListNode{0, nil}
	for cur != nil {
		tmp := ListNode{cur.Val, pre.Next}
		pre.Next = &tmp
		cur = cur.Next
	}
	return pre.Next
}
