// https://leetcode.cn/problems/remove-nth-node-from-end-of-list/?favorite=2ckc81c
package main

/**
 * Definition for singly-linked list.
 *
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

// 双指针
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{0, head}
	p, q := dummy, dummy
	tmp := p
	for i := 0; i < n; i++ {
		if q.Next != nil {
			q = q.Next
		}
	}
	for q != nil {
		tmp = p
		p = p.Next
		q = q.Next
	}
	// 循环结束，q是最后一个结点，p是要删除的结点,tmp是p前一个结点
	tmp.Next = p.Next

	return dummy.Next

}
