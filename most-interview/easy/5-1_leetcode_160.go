// https://leetcode.cn/problems/intersection-of-two-linked-lists
// 160. 相交链表

package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// 先找长度

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	p, q := headA, headB
	aLen, bLen := 0, 0
	for p != nil {
		p = p.Next
		aLen++
	}
	p = headA
	for q != nil {
		q = q.Next
		bLen++
	}
	q = headB
	if aLen >= bLen {
		for i := 0; i < aLen-bLen; i++ {
			p = p.Next
		}
	} else {
		for i := 0; i < bLen-aLen; i++ {
			q = q.Next
		}
	}
	for p != nil {
		if p == q {
			return p
		}
		p = p.Next
		q = q.Next
	}
	return p
}
