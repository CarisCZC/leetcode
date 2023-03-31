// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-ii/
package main

// 贪心解法：每天都寻求利润，所有正向利润加起来，就是最大利润
func maxProfit(prices []int) int {
	// buy := 0
	res := 0
	for i := 1; i < len(prices); i++ {
		tmp := prices[i] - prices[i-1]
		if tmp > 0 {
			res += tmp
		}
	}
	return res
}

// 动态规划
// dp[i]第i天前能获取的最大利润
// 如果prices[i]<=prices[i-1]，则dp[i] = dp[i-1]
// prices[i]>prices[i-1],则dp[i] = dp[i-1]+prices[i]-prices[i-1]
// 由于只和前一个状态有关系，因此滚动变量即可
func maxProfit2(prices []int) int {
	dp := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			dp += prices[i] - prices[i-1]
		}
	}
	return dp
}
