// https://leetcode.cn/problems/rotate-image/
// 48. 旋转图像

package main

// 先转置再翻转
func rotate(matrix [][]int) {
	for i := 0; i < len(matrix); i++ {
		for j := i; j < len(matrix[i]); j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	// 再翻转矩阵
	for i := 0; i < len(matrix); i++ {
		for j, k := 0, len(matrix)-1; j < k; {
			matrix[i][j], matrix[i][k] = matrix[i][k], matrix[i][j]
			j++
			k--
		}
	}
}
