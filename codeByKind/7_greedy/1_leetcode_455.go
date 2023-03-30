// https://leetcode.cn/problems/assign-cookies/description/
package main

import "sort"

// 这里的贪心就是：大饼干优先满足大胃口
func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	res := 0
	curs := len(s) - 1
	for i := len(g) - 1; i > -1; i-- {
		if curs < 0 {
			break
		}
		if g[i] <= s[curs] {
			res++
			curs--
		}
	}
	return res
}
