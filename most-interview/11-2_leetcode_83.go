// https://leetcode.cn/problems/remove-duplicates-from-sorted-list/
// 83. 删除排序链表中的重复元素

package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// 链表操作都要重加头结点
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	// p := &ListNode{Next: head}
	p := head.Next
	prv := head
	for p != nil {
		if p.Val == prv.Val {
			prv.Next = prv.Next.Next
		} else {
			prv = prv.Next
		}
		p = prv.Next
	}
	return head
}
