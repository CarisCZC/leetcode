package main

// 不同路径，有障碍物版本
// DP[i][j] 依旧表示到达[i][j]有多少种路径
// o[i][j] = 1表示有障碍物，此处不能走
// 初始化
// DP[0][j] = 0 or 1 // 有障碍物，即为0
// DP[i][0] = 0 OR 1
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	dp := make([][]int, m)
	dp[0] = make([]int, n)
	if obstacleGrid[0][0] == 1 {
		dp[0][0] = 0
	} else {
		dp[0][0] = 1
	}
	for i := 1; i < m; i++ {
		dp[i] = make([]int, n)
		if obstacleGrid[i][0] == 1 {
			dp[i][0] = 0 //有障碍不可达
		} else {
			dp[i][0] = dp[i-1][0]
		}

	}
	for j := 1; j < n; j++ {
		if obstacleGrid[0][j] == 1 {
			dp[0][j] = 0
		} else {
			dp[0][j] = dp[0][j-1]
		}
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				dp[i][j] = 0
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}
	return dp[m-1][n-1]
}
