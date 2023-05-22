// https://leetcode.cn/problems/merge-two-sorted-lists/
// 21. 合并两个有序链表

// Definition for singly-linked list.
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	p, q := list1, list2
	head := new(ListNode)
	cur := head
	for p != nil && q != nil {
		if p.Val <= q.Val {
			cur.Next = p
			p = p.Next
		} else {
			cur.Next = q
			q = q.Next
		}
		cur = cur.Next
	}
	if p != nil {
		cur.Next = p
	}
	if q != nil {
		cur.Next = q
	}
	return head.Next
}
