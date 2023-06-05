// https://leetcode.cn/problems/yong-liang-ge-zhan-shi-xian-dui-lie-lcof/
// 剑指 Offer 09. 用两个栈实现队列
package main

// 与https://leetcode.cn/problems/implement-queue-using-stacks/
// 同类型，但上诉题比该题更全面
type CQueue struct {
	stackIn  []int
	stackOut []int
}

func Constructor() CQueue {
	return CQueue{stackIn: []int{}, stackOut: []int{}}
}

// 尾部插入

func (this *CQueue) AppendTail(value int) {
	this.stackIn = append(this.stackIn, value)
}

func (this *CQueue) DeleteHead() int {
	if this.empty() {
		return -1
	}
	if len(this.stackOut) == 0 {
		this.stackOut = append(this.stackOut, this.stackIn...)
		this.stackIn = this.stackIn[:0]
	}
	res := this.stackOut[0]
	this.stackOut = this.stackOut[1:]
	return res
}

func (this *CQueue) empty() bool {
	return len(this.stackIn) == 0 && len(this.stackOut) == 0
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
