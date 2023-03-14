// https: //leetcode.cn/problems/remove-linked-list-elements/description/
package main

//  Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	preHead := new(ListNode)
	preHead.Next = head
	p := preHead
	for p != nil && p.Next != nil {
		if p.Next.Val == val {
			p.Next = p.Next.Next
		} else {
			p = p.Next
		}

	}
	return preHead.Next
}
