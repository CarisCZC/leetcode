// leetcode 455, 分发饼干
// 寻找第二个数组中大于等于第一个数组中的值的个数。

package main

import "sort"

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	res := 0
	i, j := len(g)-1, len(s)-1
	for i > -1 && j > -1 {
		if s[j] >= g[i] {
			res++
			j--
			i--
		} else { // 说明这个满足不了
			i--
		}
	}
	return res
}
