// https://leetcode.cn/problems/search-a-2d-matrix-ii/?favorite=2ckc81c

package main

// 矩阵从上到下升序，从左到右升序
/*
思路：1.先找每行的第一个，找到第一个比tar小的数值
2. 小了往右找，大了往上找
*/
func searchMatrix(matrix [][]int, target int) bool {
	m := 0
	n := len(matrix) - 1
	for n > -1 && m < len(matrix[0]) {
		if matrix[n][m] == target {
			return true
		}
		if matrix[n][m] < target {
			m++
		} else {
			n--
		}
	}
	return false
}
