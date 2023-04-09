package main

// 最多只能有两笔交易，且有先后顺序
// dp[i][0] 表示第i天不操作最大的利润
// dp[i][1] 表示第i天持有股票时最大的利润
// dp[i][2] 表示第i天卖出股票时最大的利润
// dp[i][3] 表示第i天第二次持有股票时最大的利润
// dp[i][4] 表示第i天第二次卖出股票时最大的利润
// 简化一下题目，可以假定，一定进行两次交易。
func maxProfit(prices []int) int {
	dp := make([][]int, len(prices))
	i := 0
	for ; i < len(prices); i++ {
		dp[i] = make([]int, 5)
	}
	dp[0][0] = 0
	// 第一次买入卖出
	dp[0][1] = -prices[0]
	dp[0][2] = 0
	// 第二次买入卖出
	dp[0][3] = -prices[0]
	dp[0][4] = 0
	for i := 1; i < len(prices); i++ {
		dp[i][0] = dp[i-1][0]
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
		dp[i][2] = max(dp[i-1][2], dp[i-1][1]+prices[i])
		dp[i][3] = max(dp[i-1][3], dp[i-1][2]-prices[i])
		dp[i][4] = max(dp[i-1][4], dp[i-1][3]+prices[i])
	}
	return dp[i-1][4]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// dp分两段
func maxProfit2(prices []int) int {
	dp := make([]int, len(prices))
	// 从第0天到第i天的最大利润
	dp[0] = 0
	i := 1
	curPrices := prices[0]
	for ; i < len(prices); i++ {
		if prices[i] > curPrices {
			tmp := prices[i] - curPrices
			if tmp > dp[i-1] {
				dp[i] = tmp
			} else {
				dp[i] = dp[i-1]
			}
		} else {
			dp[i] = dp[i-1]
			curPrices = prices[i]
		}
	}
	// 从第i天到最后一天的最大利润
	dp2 := make([]int, i)
	dp2[i-1] = 0
	curPrices = prices[i-1]
	i -= 2
	for ; i >= 0; i-- {
		if prices[i] < curPrices {
			tmp := curPrices - prices[i]
			if tmp > dp2[i] {
				dp2[i] = tmp
			} else {
				dp2[i] = dp2[i+1]
			}
		} else {
			dp2[i] = dp2[i+1]
			curPrices = prices[i]
		}
	}
	res := dp[0] + dp2[0]
	for j := 1; j < len(prices); j++ {
		if dp[j]+dp2[j] > res {
			res = dp[j] + dp2[j]
		}
	}
	return res
}

func main() {
	maxProfit2([]int{3, 3, 5, 0, 0, 3, 1, 4})
}
