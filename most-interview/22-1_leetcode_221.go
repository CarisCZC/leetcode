// https://leetcode.cn/problems/maximal-square/
// 221. 最大正方形

package main

// dp[i][j]表示以[i,j]为右下角的正方形的最大边长
// dp[i][j]取决于dp[i-1][j] dp[i][j-1] dp[i-1][j-1]
func maximalSquare(matrix [][]byte) int {
	dp := make([][]int, len(matrix))
	for i := range dp {
		dp[i] = make([]int, len(matrix[i]))
	}
	res := 0
	// 初始化第一行第一列
	for i := range matrix {
		dp[i][0] = int(matrix[i][0] - '0')
		if dp[i][0] > res {
			res = dp[i][0]
		}
	}
	for i := 0; i < len(matrix[0]); i++ {
		dp[0][i] = int(matrix[0][i] - '0')
		if dp[0][i] > res {
			res = dp[0][i]
		}
	}
	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[i]); j++ {
			if matrix[i][j] == '0' {
				dp[i][j] = int(matrix[i][j] - '0')
			} else {
				dp[i][j] = min(dp[i-1][j], min(dp[i-1][j-1], dp[i][j-1])) + 1
			}
			if dp[i][j] > res {
				res = dp[i][j]
			}
		}

	}
	return res * res
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
