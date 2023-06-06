// https://leetcode.cn/problems/reorder-list/
// 143. 重排链表
package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// 逆置后半链表，然后进行插入
func reorderList(head *ListNode) {
	fast, slow := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	// slow会停在需要逆转的前一个结点
	halfOfEnd := reverse(slow.Next)
	slow.Next = nil
	cur := head
	for halfOfEnd != nil && cur != nil {
		tmp := halfOfEnd
		next := cur.Next
		halfOfEnd = halfOfEnd.Next
		cur.Next = tmp
		tmp.Next = next
		cur = next
	}
}

func reverse(head *ListNode) *ListNode {
	p := &ListNode{}
	for head != nil {
		tmp := head
		head = head.Next
		next := p.Next
		p.Next = tmp
		tmp.Next = next
	}
	return p.Next
}
