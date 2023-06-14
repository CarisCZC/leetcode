// https://leetcode.cn/problems/implement-rand10-using-rand7/
// 470. 用 Rand7() 实现 Rand10()

package main

func rand10() int {
	for {
		row := rand7()
		col := rand7()
		idx := (row-1)*7 + col
		if idx <= 40 {
			return 1 + (idx-1)%10
		}
	}
}
