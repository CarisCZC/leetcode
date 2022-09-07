// leetcode 657, 机器人能否返回原点

// 步伐抵消: 一左一右抵消，一前一后抵消

package main

func judgeCircle(moves string) bool {
	ucnt, lcnt := 0, 0
	for _, v := range moves {
		if v == 'U' {
			ucnt++
		}
		if v == 'D' {
			ucnt--
		}
		if v == 'L' {
			lcnt++
		}
		if v == 'R' {
			lcnt--
		}
	}
	if ucnt == 0 && lcnt == 0 {
		return true
	}
	return false
}
