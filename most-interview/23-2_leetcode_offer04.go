package main

// 找中线
// 比当前数大，但比下一个数小，找该列下一个，该行下一个
func findNumberIn2DArray(matrix [][]int, target int) bool {
	i, j := len(matrix)-1, 0
	for i > 0 && j < len(matrix[0]) {
		if matrix[i][j] == target {
			return true
		}
		if matrix[i][j] > target {
			i--

		} else {
			j++
		}
	}
	return false
}
