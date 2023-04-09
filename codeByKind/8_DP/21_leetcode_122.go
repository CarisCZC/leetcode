package main

// 多买多卖

// dp[i]表示在第i天可以获得的最大利润
// dp[0] = 0
// 如果prices[i]<prices[i-1],则dp = dp[i-1]
// 否则，dp = dp[i-1]+prices[i]-prices[i-1]
func maxProfit(prices []int) int {
	dp := make([]int, len(prices))
	dp[0] = 0
	for i := 1; i < len(prices); i++ {
		if dp[i] > dp[i-1] {
			dp[i] = dp[i-1] + prices[i] - prices[i-1]
		} else {
			dp[i] = dp[i-1]
		}
	}
	return dp[len(prices)-1]
}
