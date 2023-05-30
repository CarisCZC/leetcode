// https://leetcode.cn/problems/sqrtx/
// 69. x 的平方根

package main

// 实际上，考查的是二分查找
func mySqrt(x int) int {
	y := 0
	cur := x
	for y <= cur {
		y++
		cur = x
		cur /= y
	}
	return y - 1
}
