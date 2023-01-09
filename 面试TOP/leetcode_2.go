// https://leetcode.cn/problems/add-two-numbers/?favorite=2ckc81c
// 两数相加
package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	res := ListNode{0, nil}
	p := &res
	plus := 0
	for l1 != nil || l2 != nil {
		tmp := plus
		plus = 0
		if l1 != nil {
			tmp += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			tmp += l2.Val
			l2 = l2.Next
		}
		if tmp >= 10 {
			plus = 1
			tmp -= 10
		}
		tmpNode := ListNode{tmp, nil}
		p.Next = &tmpNode
		p = p.Next
	}
	if plus != 0 {
		p.Next = &ListNode{plus, nil}
	}
	return res.Next
}
