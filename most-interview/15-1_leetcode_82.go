// https://leetcode.cn/problems/remove-duplicates-from-sorted-list-ii/
// 82. 删除排序链表中的重复元素 II
package main

//   Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// 双指针
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	p := &ListNode{Next: head}
	cur := p
	for cur.Next != nil && cur.Next.Next != nil {
		if cur.Next.Val == cur.Next.Next.Val {
			x := cur.Next.Val
			for cur.Next != nil && cur.Next.Val == x {
				cur.Next = cur.Next.Next
			}
		} else {
			cur = cur.Next
		}
	}
	return p.Next
}
