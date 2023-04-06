package main

// 数组定义 dp[i][j] 走到[i][j]的最小路径和
// 递推公式 dp[i][j] = min(dp[i-1][j]+grid[i][j],dp[i][j-1]+grid[i][j])
// 初始化 dp[0][0] = 0
// 循环遍历
func minPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	dp[0][0] = grid[0][0]
	for j := 1; j < n; j++ {
		dp[0][j] = dp[0][j-1] + grid[0][j]
	}
	for i := 1; i < m; i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = min(dp[i-1][j]+grid[i][j], dp[i][j-1]+grid[i][j])
		}
	}
	return dp[m-1][n-1]
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
