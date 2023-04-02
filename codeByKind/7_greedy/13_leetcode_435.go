package main

import (
	"sort"
)

// 与前一道题很相似
// 贪心贪的就是，最小的end
func eraseOverlapIntervals(intervals [][]int) int {
	// 先排序一下
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0] || intervals[i][0] == intervals[j][0] && intervals[i][1] < intervals[j][1]
	})
	end := intervals[0][1] // 第一个开始
	res := 0
	for i := 1; i < len(intervals); i++ { // 能碰出到的区间，都得删除
		if end > intervals[i][0] {
			// 删除，但应该保留当前有最小的end的那个区间
			end = min(end, intervals[i][1])
			res++
		} else { // 说明当前end区间与i不相交。查看i是否与后面的区间相交
			end = intervals[i][1]
		}
	}
	return res
}
func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
