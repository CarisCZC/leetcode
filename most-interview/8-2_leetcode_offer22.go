// https://leetcode.cn/problems/lian-biao-zhong-dao-shu-di-kge-jie-dian-lcof/

// 剑指 Offer 22. 链表中倒数第k个节点
package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// 双指针
func getKthFromEnd(head *ListNode, k int) *ListNode {
	p, q := head, head
	for i := 0; i < k; i++ {
		if q != nil {
			q = q.Next
		} else {
			return nil
		}
	}
	for q != nil {
		p = p.Next
		q = q.Next
	}
	return p
}
