// leetcode 766 托普利茨矩阵

package main

// 对角线的元素都相等
func isToeplitzMatrix(matrix [][]int) bool {

	// 先横着查对角线
	for i := 0; i < len(matrix[0]); i++ {
		row, col := 0, i
		for row < len(matrix)-1 && col < len(matrix[0])-1 {
			if matrix[row][col] == matrix[row+1][col+1] {
				row++
				col++
			} else {
				return false
			}
		}
	}
	//竖着对角线查
	for i := 1; i < len(matrix); i++ {
		row, col := i, 0
		for row < len(matrix)-1 && col < len(matrix[0])-1 {
			if matrix[row][col] == matrix[row+1][col+1] {
				row++
				col++
			} else {
				return false
			}
		}
	}
	return true
}
