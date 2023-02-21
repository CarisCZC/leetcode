// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-ii/?favorite=2ckc81c

package main

// 所有的升序片段
func maxProfit(prices []int) int {
	res := 0
	buy := 0
	for i := 1; i < len(prices); i++ {
		if prices[buy] < prices[i] {
			res += prices[i] - prices[buy]
		}
		buy = i
	}
	return res
}
