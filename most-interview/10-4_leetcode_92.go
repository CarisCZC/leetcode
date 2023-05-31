// https://leetcode.cn/problems/reverse-linked-list-ii/
// 92. 反转链表 II

package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// 记录left的prv，right的next。中间翻转完成后，重新连接
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	cnt := 0
	q := &ListNode{Next: head}
	leftPrv := q
	p := q
	tail := &ListNode{}
	for cnt <= right {
		if cnt == left-1 {
			leftPrv = p
		}
		if cnt == left {
			tail = p
		}
		p = p.Next
		cnt++
	}
	leftPrv.Next = reverse(leftPrv.Next, left, right)
	tail.Next = p
	return q.Next
}

func reverse(head *ListNode, left int, right int) *ListNode {
	p := &ListNode{}
	q := head
	// tail := &ListNode{}
	for q != nil && left <= right {
		next := p.Next
		tmp := q
		q = q.Next
		p.Next = tmp
		tmp.Next = next
		left++
	}
	return p.Next
}
