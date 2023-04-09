package main

// 卖股票的后一天是冷冻期
// dp[i] 表示第i天能获得的最大利润
// dp[i][2] = dp[i-1][2] //如果i-1天卖了股票
// dp[i][2] = dp[i-1][1]+prices[i]
// dp[i][1] = dp[i-1][0] //第i-1天卖了股票 也就是当天不能操作
// dp[i][1] = dp[i-1][0]-prices[i]
func maxProfit(prices []int) int {
	n := len(prices)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, 4)
	}
	dp[0][0] = -prices[0] // 当天持有股票
	dp[0][1] = 0          // 保持卖出状态
	dp[0][2] = 0          // 当天卖出股票
	dp[0][3] = dp[0][2]   // 当天为冷冻期时，最大利润。也就是dp[i][3] = dp[i-1][2]
	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], max(dp[i-1][3]-prices[i], dp[i-1][1]-prices[i]))
		dp[i][1] = max(dp[i-1][1], dp[i-1][3])
		dp[i][2] = dp[i-1][0] + prices[i]
		dp[i][3] = dp[i-1][2]

	}
	return max(dp[n-1][3], max(dp[n-1][1], dp[n-1][2]))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
