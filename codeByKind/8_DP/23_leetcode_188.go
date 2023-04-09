// 就是123题的扩展版本
package main

func maxProfit(k int, prices []int) int {
	dp := make([][]int, len(prices))
	for i := range dp {
		dp[i] = make([]int, k*2+1)
	}
	for i := 0; i < k*2+1; i++ {
		if i%2 == 0 {
			dp[0][i] = 0
		} else {
			dp[0][i] = -prices[0]
		}
	}
	i := 1
	for ; i < len(prices); i++ {
		dp[i][0] = dp[i-1][0]
		for j := 1; j < k*2+1; j++ {
			if j%2 != 0 {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-1]-prices[i])
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-1]+prices[i])
			}
		}
	}
	return dp[i-1][2*k]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	maxProfit(4, []int{1, 2, 4, 2, 5, 7, 2, 4, 9, 0})
}
