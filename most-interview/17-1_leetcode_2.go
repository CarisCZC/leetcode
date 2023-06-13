// https://leetcode.cn/problems/add-two-numbers/
// 2. 两数相加

package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	upset := 0
	head1, head2 := l1, l2
	dummyHead := &ListNode{}
	res := dummyHead
	for head1 != nil && head2 != nil {
		cur := head1.Val + head2.Val + upset
		upset = cur / 10
		cur = cur % 10
		res.Next = &ListNode{cur, nil}
		res = res.Next
		head1 = head1.Next
		head2 = head2.Next
	}
	for head1 != nil {
		cur := head1.Val + upset
		upset = cur / 10
		cur = cur % 10
		res.Next = &ListNode{cur, nil}
		res = res.Next
		head1 = head1.Next
	}
	for head2 != nil {
		cur := head2.Val + upset
		upset = cur / 10
		cur = cur % 10
		res.Next = &ListNode{cur, nil}
		res = res.Next
		head2 = head2.Next
	}
	if upset != 0 {
		res.Next = &ListNode{upset, nil}
	}
	return dummyHead.Next
}
