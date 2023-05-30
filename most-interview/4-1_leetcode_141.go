// https://leetcode.cn/problems/linked-list-cycle/
// 141. 环形链表
// 判断链表中是否有环

package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// 双指针
func hasCycle(head *ListNode) bool {
	slow, fast := head, head
	for slow != nil && fast != nil {
		if slow != nil {
			slow = slow.Next
		}
		for i := 0; i < 2; i++ {
			if fast != nil {
				fast = fast.Next
			}
		}
		if fast == slow && fast != nil {
			return true
		}
	}
	return false
}
