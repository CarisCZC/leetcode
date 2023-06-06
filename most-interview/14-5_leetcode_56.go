// https://leetcode.cn/problems/merge-intervals/
// 56. 合并区间

package main

import "sort"

// 先对区间进行排序，然后合并
func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0] || (intervals[i][0] == intervals[j][0] && intervals[i][1] < intervals[j][1])
	})

	res := [][]int{}
	perv := intervals[0]
	for i := 1; i < len(intervals); i++ {
		cur := intervals[i]
		if perv[1] >= cur[0] {
			perv[1] = max(perv[1], cur[1])
		} else {
			res = append(res, perv)
			perv = cur
		}
	}
	res = append(res, perv)
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
