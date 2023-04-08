package main

// dp[i][0] 表示第i天拥有股票时的最大现金
// dp[i][1] 表示第i天不拥有股票的最大现金
// dp[i][0] = max(dp[i-1][0],-prices[i])
// dp[i][1] = max(dp[i-1][1],prices[i]+dp[i-1][0])

// dp[0][0] = -prices[i]
// dp[0][1] = 0
func maxProfit(prices []int) int {
	dp := make([][]int, len(prices))
	for i := range dp {
		if i == 0 {
			dp[i] = []int{-prices[i], 0}
		} else {
			dp[i] = []int{0, 0}
		}

	}
	for i := 1; i < len(prices); i++ {
		dp[i] = []int{max(dp[i-1][0], -prices[i]), max(dp[i-1][1], prices[i]+dp[i-1][0])}
	}
	return dp[len(prices)-1][1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
