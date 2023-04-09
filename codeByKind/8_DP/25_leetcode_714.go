package main

// 如果当天卖出，需要收一笔手续费
// 三个状态
// 1. 持有股票的时候
// 2. 保持卖出状态
// 3. 当天卖出时的收益
func maxProfit(prices []int, fee int) int {

	n := len(prices)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, 3)
	}
	dp[0][0] = -prices[0]
	dp[0][1] = 0
	dp[0][2] = -fee
	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], max(dp[i-1][1]-prices[i], dp[i-1][2]-prices[i]))
		dp[i][1] = max(dp[i-1][1], dp[i-1][2])
		dp[i][2] = dp[i-1][0] + prices[i] - fee
	}
	return max(dp[n-1][1], dp[n-1][2])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	maxProfit([]int{1, 3, 2, 8, 4, 9}, 2)
}
