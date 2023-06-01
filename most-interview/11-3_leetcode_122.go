// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-ii/
// 122. 买卖股票的最佳时机 II
package main

// dp或者贪心，都行
// dp[i] 表示第i天的最大利润
// dp[i] 可以简化成dp2
func maxProfit(prices []int) int {
	// dp := make([]int, len(prices))
	// dp[0] = 0
	dp2 := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			dp2 += prices[i] - prices[i-1]
			// dp[i] = dp[i-1] + prices[i] - prices[i-1]
		} else {
			// dp[i] = dp[i-1]
		}
	}
	// return dp[len(prices)-1]
	return dp2
}
