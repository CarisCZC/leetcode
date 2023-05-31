// https://leetcode.cn/problems/linked-list-cycle-ii/description/
// 142. 环形链表 II

package main

//   Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// 假设链表入口长度a，环长度为b
//  f = 2s fast走的步数是s的两倍
// 而f = s + nb  即fast 比 slow 多走了n圈
// 可得 s = nb。而达到入口 k=a+nb 可知，s再走a步一定到入口
// 而head 走a步也到入口，故有p = s，此时s就停在入口处
func detectCycle(head *ListNode) *ListNode {
	fast, slow := head, head
	for {
		if fast == nil || fast.Next == nil {
			return nil //无环
		}
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow { // 第一次相遇
			break
		}
	}
	fast = head
	for {
		fast = fast.Next
		slow = slow.Next
		if fast == slow {
			break
		}
	}
	return slow
}
