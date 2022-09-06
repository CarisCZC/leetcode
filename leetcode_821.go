// leetcode 821,字符的最短距离

package main

import "math"

// 记录c的坐标
func shortestToChar(s string, c byte) []int {
	res := make([]int, len(s))
	for i := 0; i < len(s); i++ {
		res[i] = math.MaxInt
	}
	indexs := []int{}
	for i := 0; i < len(s); i++ {
		if s[i] == c {
			indexs = append(indexs, i)
		}
	}
	for i := 0; i < len(indexs); i++ {
		for j := 0; j < len(res); j++ {
			res[j] = min(int(math.Abs(float64(j-indexs[i]))), res[j])
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
