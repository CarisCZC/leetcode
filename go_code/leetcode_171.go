// leetcode 171. Excel 表列序号

// 和168题是相反的函数操作

package main

func titleToNumber(columnTitle string) int {
	res := 0
	base := 1
	for i := len(columnTitle) - 1; i > -1; i-- {
		res += base * (int(columnTitle[i]) - 64)
		base *= 26
	}
	return res
}

