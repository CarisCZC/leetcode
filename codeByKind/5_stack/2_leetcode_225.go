package main

type MyStack struct {
	stack []int
	len   int
}

func Constructor() MyStack {
	return MyStack{stack: make([]int, 0), len: 0}
}

// 入栈
func (this *MyStack) Push(x int) {
	this.stack = append(this.stack, x)
	this.len++
}

// 出栈
func (this *MyStack) Pop() int {
	if this.len == 0 {
		return -1
	}
	val := this.stack[this.len-1]
	this.stack = this.stack[:this.len-1]
	this.len--
	return val
}

// 栈顶元素
func (this *MyStack) Top() int {
	if this.len == 0 {
		return -1
	}
	return this.stack[this.len-1]
}

func (this *MyStack) Empty() bool {
	return this.len == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
