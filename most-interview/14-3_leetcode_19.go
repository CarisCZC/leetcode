// https://leetcode.cn/problems/remove-nth-node-from-end-of-list/description/
// 19. 删除链表的倒数第 N 个结点

// Definition for singly-linked list.
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 双指针
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	p := &ListNode{Next: head}
	fast := p
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	slow := p
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	// 停止位置：删除结点的前一个
	slow.Next = slow.Next.Next
	return p.Next
}
