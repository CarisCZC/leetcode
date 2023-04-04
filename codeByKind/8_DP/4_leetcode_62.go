package main

// 这题如果看作是数学问题就是排列组合问题
// C（m-1）（m+n-2）
// 要走m-1步向下和n-1步向左
// 动态规划
// 初始：i=1，j=1
// 结束：i=m,j=n
// 每次：i+1或者j+1
// DP[i][j] 表示到达i，j有多少条路径
// DP[i][j] = DP[i-1][j]+DP[i][j-1]
// DP[0][i] = 1
// DP[i][0] = 1
func uniquePaths(m int, n int) int {
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
		dp[i][0] = 1
	}
	for j := 0; j < n; j++ {
		dp[0][j] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}
