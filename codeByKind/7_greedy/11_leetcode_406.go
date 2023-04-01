// https: //leetcode.cn/problems/queue-reconstruction-by-height/

package main

import "sort"

// 怎么感觉，这更多的是一个排序？
func reconstructQueue(people [][]int) [][]int {
	// 先排序
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] == people[j][0] {
			return people[i][1] < people[j][1]
		}
		return people[i][0] > people[j][0]
	})
	res := make([][]int, 0, len(people))
	for i := 0; i < len(people); i++ {
		spaces := people[i][1]
		res = append(res[:spaces], append([][]int{people[i]}, res[spaces:]...)...)
	}
	return res
}
