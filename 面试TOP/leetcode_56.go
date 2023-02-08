// https://leetcode.cn/problems/merge-intervals/?favorite=2ckc81c
package main

import "sort"

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	res := [][]int{}
	for _, interval := range intervals {
		length := len(res)
		// 按递增排序后，扩展的区间一定是res最后一个区间
		if length == 0 || res[length-1][1] < interval[0] {
			res = append(res, interval)
		} else {
			res[length-1][1] = max(res[length-1][1], interval[1])
		}
	}
	return res
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
