package main

type MyQueue struct {
	stackIn  []int
	stackOut []int
}

func Constructor() MyQueue {
	return MyQueue{
		stackIn:  make([]int, 0),
		stackOut: make([]int, 0),
	}
}

// 入队
func (this *MyQueue) Push(x int) {
	this.stackIn = append(this.stackIn, x)
}

// 出队
func (this *MyQueue) Pop() int {
	if this.Empty() {
		return -1
	}
	inlen, outlen := len(this.stackIn), len(this.stackOut)
	if outlen == 0 {
		for inlen != 0 {
			this.stackOut = append(this.stackOut, this.stackIn[inlen-1])
			inlen--
		}
		//此时inlen=0
		this.stackIn = this.stackIn[:inlen]
	}
	outlen = len(this.stackOut)
	value := this.stackOut[outlen-1]
	this.stackOut = this.stackOut[:outlen-1]
	return value
}

// 返回队列开头的元素
func (this *MyQueue) Peek() int {
	val := this.Pop()
	if val != -1 {
		this.stackOut = append(this.stackOut, val)
	}
	return val
}

// 是否为空
func (this *MyQueue) Empty() bool {
	return len(this.stackIn) == 0 && len(this.stackOut) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
