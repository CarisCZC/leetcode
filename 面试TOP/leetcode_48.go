// https://leetcode.cn/problems/rotate-image/?favorite=2ckc81c
package main

// 旋转图像，第i行变成第n-i列
func rotate(matrix [][]int) {
	n := len(matrix)
	// 先行列变换
	for i := 0; i < n; i++ {
		// 取行
		for j := i; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	// 然后每行翻转
	for _, row := range matrix {
		for i := 0; i < n/2; i++ {
			row[i], row[n-i-1] = row[n-i-1], row[i]
		}
	}
}

func main() {
	rotate([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}})
}
