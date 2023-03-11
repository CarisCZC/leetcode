// https://leetcode.cn/problems/intersection-of-two-linked-lists-lcci/description/
// 面试题02.07
package main

//  Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// 给你两个单链表的头节点 headA 和 headB ，请你找出并返回两个单链表相交的起始节点。如果两个链表没有交点，返回 null

// 末端对其
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	// 1. 找出A,B的长度
	lenA, lenB := findLen(headA), findLen(headB)
	// 长的先走
	curA, curB := headA, headB
	if lenA >= lenB {
		tmp := lenA - lenB
		for tmp > 0 {
			curA = curA.Next
			tmp--
		}
	} else {
		tmp := lenB - lenA
		for tmp > 0 {
			curB = curB.Next
			tmp--
		}
	}
	// 一起走
	// curA != nil && curB != nil &&
	for curA != curB {
		curA = curA.Next
		curB = curB.Next
	}
	return curA
}

func findLen(head *ListNode) int {
	sum := 0
	for head != nil {
		head = head.Next
		sum++
	}
	return sum
}
