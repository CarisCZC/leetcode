// https://leetcode.cn/problems/minimum-path-sum/
// 64. 最小路径和

package main

// 从00 走到 ij
// dp[i][j] 表示走到 i，j的最小路径和
// dp[0][0] = grid[0][0]
// dp[i][j] = min(dp[i][j-1],dp[i-1][j])+dp[i][j]
func minPathSum(grid [][]int) int {
	dp := make([][]int, len(grid))
	for i := range dp {
		dp[i] = make([]int, len(grid[i]))
	}
	dp[0][0] = grid[0][0]
	// 第一行第一列得先初始化
	for i := 1; i < len(grid); i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}
	for j := 1; j < len(grid[0]); j++ {
		dp[0][j] = dp[0][j-1] + grid[0][j]
	}
	for i := 1; i < len(grid); i++ {
		for j := 1; j < len(grid[i]); j++ {
			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}
	return dp[len(grid)-1][len(grid[0])-1]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	minPathSum([][]int{{1, 2, 3}, {4, 5, 6}})
}
