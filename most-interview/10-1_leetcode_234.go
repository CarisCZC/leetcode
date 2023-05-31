// https://leetcode.cn/problems/palindrome-linked-list/
// 234. 回文链表
package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// 判断链表是回文串
// 1. 放入数组判断
// 2. 递归，这个还不如1
// 3. 改变输入，逆转后半段
func isPalindrome(head *ListNode) bool {
	fast, slow := head, head
	for fast != nil {
		if fast.Next != nil {
			fast = fast.Next.Next
		} else {
			fast = fast.Next
		}
		slow = slow.Next
	}
	slow = traverse(slow)
	fast = head
	for slow != nil {
		if slow.Val != fast.Val {
			return false
		}
		slow = slow.Next
		fast = fast.Next
	}
	return true
}

func traverse(head *ListNode) *ListNode {
	p := *&ListNode{}
	for head != nil {
		tmp := head
		head = head.Next
		next := p.Next
		p.Next = tmp
		tmp.Next = next
	}
	return p.Next
}
