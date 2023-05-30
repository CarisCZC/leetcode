// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock/
// 121. 买卖股票的最佳时机

package main

// Dp思想，不过这题是简单Dp，用一个变量即可
func maxProfit(prices []int) int {
	minPrice := prices[0]
	cnt := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else if prices[i]-minPrice > cnt {
			cnt = prices[i] - minPrice
		}
	}
	return cnt
}
