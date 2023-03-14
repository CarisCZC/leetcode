package main

// 设置一个虚拟头结点，不要在意内存消耗，重点要方便写代码
type ListNode struct {
	Val  int
	Next *ListNode
}
type MyLinkedList struct {
	head *ListNode
	size int
}

func Constructor() MyLinkedList {
	return MyLinkedList{&ListNode{}, 0}
}

func (this *MyLinkedList) Get(index int) int {
	if index < 0 || index >= this.size {
		return -1
	}
	cur := this.head
	for i := 0; i <= index; i++ {
		cur = cur.Next
	}
	return cur.Val

}

func (this *MyLinkedList) AddAtHead(val int) {
	this.AddAtIndex(0, val)
}

func (this *MyLinkedList) AddAtTail(val int) {
	this.AddAtIndex(this.size, val)
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index > this.size {
		return
	}
	index = max(index, 0)
	pred := this.head
	for i := 0; i < index; i++ {
		pred = pred.Next
	}
	toAdd := ListNode{val, pred.Next}
	pred.Next = &toAdd
	this.size++
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index > -1 && index < this.size {
		cur := this.head
		for i := 0; i < index; i++ {
			cur = cur.Next
		}
		cur.Next = cur.Next.Next
		this.size--
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
