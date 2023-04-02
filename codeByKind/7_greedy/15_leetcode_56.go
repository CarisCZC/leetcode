package main

import "sort"

// 这个题与12，13都比较类似
func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0] || intervals[i][0] == intervals[j][0] && intervals[i][1] < intervals[j][1]
	})
	res := [][]int{intervals[0]}
	cur := 0
	for i := 1; i < len(intervals); i++ {
		if res[cur][1] >= intervals[i][0] { // 保留最大的end
			res[cur][1] = max(intervals[i][1], res[cur][1])
		} else {
			res = append(res, intervals[i])
			cur++
		}
	}
	return res
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
