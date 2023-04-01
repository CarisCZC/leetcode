package main

import "sort"

// 尽可能的找相交
func findMinArrowShots(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0] || points[i][0] == points[j][0] && points[i][1] < points[j][1]
	})
	sameArea := [][]int{points[0]}
	// 排序后，这个就不必要了
	for i := 1; i < len(points); i++ {
		canSame := false
		cur := points[i]
		for j := 0; j < len(sameArea); j++ {
			newStart, newEnd := max(cur[0], sameArea[j][0]), min(cur[1], sameArea[j][1])
			if newEnd >= newStart { //相交
				sameArea[j] = []int{newStart, newEnd}
				canSame = true
				break
			}
		}
		if !canSame {
			sameArea = append(sameArea, cur)
			canSame = false
		}
	}
	return len(sameArea)
}

func findMinArrowShots2(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0] || points[i][0] == points[j][0] && points[i][1] < points[j][1]
	})
	// }
	// 换个写法
	// 排序后，左边的要和右边的相交，一定有leftEnd >= rightStart
	curEnd := points[0][1]
	res := 1
	for i := 0; i < len(points); i++ {
		if curEnd < points[i][0] { // 说明不相交
			res++                 //多射一箭
			curEnd = points[i][1] // 换新的交点尝试
		} else { //相交了也要改变一下curEnd
			curEnd = min(curEnd, points[i][1])
		}
	}
	return res
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
