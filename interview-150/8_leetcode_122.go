// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-ii/?envType=study-plan-v2&id=top-interview-150
// 122. 买卖股票的最佳时机 II
package main

// 可动态规划
// 也可以求每天和下一天的顺差 就是贪心。
// 两者思路都类似
func maxProfit(prices []int) int {
	dp := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			dp += prices[i] - prices[i-1]
		}
	}
	return dp
}
