package main

// dp[i][j] 表示下标范围[i,j]中，当前玩家与另一个玩家分数之差的最大值
// i>j没有意义，i=j时dp[i][j] = nums[i]
// i<j时 dp[i][j] = max(nums[i]-dp[i+1][j],nums[j]-dp[i][j-1])
func PredictTheWinner(nums []int) bool {
	n := len(nums)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
		dp[i][i] = nums[i]
	}
	for i := n - 1; i >= 0; i-- { // i=n-1时不进入循环，无意义，但好理解
		for j := i + 1; j < n; j++ {
			// 当前玩家选择i或者j，那么下一个玩家的选择就是dp[i+1][j]，或者dp[i][j-1]
			// 此时就出现了差值
			dp[i][j] = max(nums[i]-dp[i+1][j], nums[j]-dp[i][j-1])
		}
	}
	return dp[0][n-1] >= 0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
