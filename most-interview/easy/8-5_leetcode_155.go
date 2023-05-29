// https://leetcode.cn/problems/min-stack/
// 155. 最小栈
package main

import "math"

type MinStack struct {
	stack    []int
	minstack []int
}

func Constructor() MinStack {
	return MinStack{
		stack:    []int{},
		minstack: []int{math.MaxInt64},
	}
}

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)
	top := this.minstack[len(this.minstack)-1]
	this.minstack = append(this.minstack, min(val, top))

}

func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
	this.minstack = this.minstack[:len(this.minstack)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.minstack[len(this.minstack)-1]
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
