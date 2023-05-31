// https://leetcode.cn/problems/spiral-matrix/
// 54. 螺旋矩阵
package main

// 维护四个边界
func spiralOrder(matrix [][]int) []int {
	res := []int{}
	top, bottom, left, right := 0, len(matrix)-1, 0, len(matrix[0])-1
	// 一圈圈的循环
	for top <= bottom && left <= right {
		// 上边界
		for j := left; j <= right; j++ {
			res = append(res, matrix[top][j])
		}
		// 右边界
		for i := top + 1; i <= bottom; i++ {
			res = append(res, matrix[i][right])
		}
		// 下边界
		if top < bottom && left < right { // 保证top和bottom不是同一行。
			for j := right - 1; j >= left; j-- {
				res = append(res, matrix[bottom][j])
			}
			//左边界
			for i := bottom - 1; i > top; i-- {
				res = append(res, matrix[i][left])
			}
		}
		top++
		bottom--
		left++
		right--
	}
	return res
}
