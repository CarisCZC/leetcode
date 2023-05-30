package main

// go没栈，就用切片了
type MyQueue struct {
	stackIn  []int
	stackOut []int
}

func Constructor() MyQueue {
	return MyQueue{stackIn: []int{}, stackOut: []int{}}
}

// 将元素加入队尾
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
			outlen++
		}
		this.stackIn = this.stackIn[:inlen] // 清空入队元素
	}
	value := this.stackOut[outlen-1]
	this.stackOut = this.stackOut[:outlen-1]
	return value
}

func (this *MyQueue) Peek() int {
	val := this.Pop()
	if val != -1 {
		this.stackOut = append(this.stackOut, val)
	}
	return val
}

func (this *MyQueue) Empty() bool {
	return len(this.stackIn) == 0 && len(this.stackOut) == 0
}
