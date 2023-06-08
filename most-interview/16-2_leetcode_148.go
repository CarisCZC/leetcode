// https://leetcode.cn/problems/sort-list/
// 148. 排序链表

package main

//   Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	listLen := 0
	p := head
	for p != nil {
		listLen++
		p = p.Next
	}
	dummyHead := &ListNode{Next: head}
	for subLength := 1; subLength < listLen; subLength <<= 1 {
		prev, cur := dummyHead, dummyHead.Next
		for cur != nil {
			// 第一个subLength长度的链
			head1 := cur
			for i := 1; i < subLength && cur.Next != nil; i++ {
				cur = cur.Next
			}
			// 第二个subLength长度的链
			head2 := cur.Next
			// 这里先断掉head1最后一个元素
			// 否则会引起链表嵌套成环
			cur.Next = nil
			cur = head2
			for i := 1; i < subLength && cur != nil && cur.Next != nil; i++ {
				cur = cur.Next
			}
			// 这里同理，找到第二个subLength长度链的最后一个元素，然后进行断链操作
			var next *ListNode
			if cur != nil {
				next = cur.Next
				cur.Next = nil
			}
			prev.Next = merge(head1, head2)
			// prev到达dummyHead最后一个结点。用于下一次链的拼接
			for prev.Next != nil {
				prev = prev.Next
			}
			cur = next
		}
	}
	return dummyHead.Next
}

// 合并两个有序链表
func merge(head1, head2 *ListNode) *ListNode {
	dummyHead := &ListNode{}
	temp, temp1, temp2 := dummyHead, head1, head2
	for temp1 != nil && temp2 != nil {
		if temp1.Val <= temp2.Val {
			temp.Next = temp1
			temp1 = temp1.Next
		} else {
			temp.Next = temp2
			temp2 = temp2.Next
		}
		temp = temp.Next
	}
	if temp1 != nil {
		temp.Next = temp1
	}
	if temp2 != nil {
		temp.Next = temp2
	}
	return dummyHead.Next
}
