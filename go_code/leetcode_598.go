// leetcode 598, 范围求和Ⅱ

package main

func maxCount(m int, n int, ops [][]int) int {
	row := make([]int, m)
	col := make([]int, n)
	for _, op := range ops {
		row[op[0]-1]++
		col[op[1]-1]++
	}
	ansr := 0
	for i := 0; i < m; i++ {
		if row[i] != 0 {
			ansr = i + 1
			break
		}
		if i == m-1 {
			ansr = m
		}
	}
	ansc := 0
	for i := 1; i < n; i++ { //第一个不为0的数字
		if col[i] != 0 {
			ansc = i + 1
			break
		}
		if i == m-1 {
			ansc = n
		}
	}
	return ansr * ansc
}

// 大致思路对，但是没有官方案例精简
// 案例意思是，如果op中有个[3,3]，那么小于3的都会被操作（即前两行前两列）。于是我们只需要找最小的行和列就好
// 我的思路是，将行和列操作存储起来，最终会有个行列的数组，然后找第一个非0的行。（后面跟行列的操作都会累加在前面的行列上）
