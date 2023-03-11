package main

//  Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// 双指针法
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{0, head}
	cur, re := dummy, dummy
	tmp := cur
	for i := 0; i < n; i++ {
		re = re.Next
	}
	for re != nil {
		tmp = cur
		re = re.Next
		cur = cur.Next
	}
	// 循环结束，cur就是要被删除的节点，tmp是其前一个结点
	tmp.Next = cur.Next
	return dummy.Next

}
