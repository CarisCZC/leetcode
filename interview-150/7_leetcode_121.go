// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock/?envType=study-plan-v2&id=top-interview-150
// 买卖股票的最佳时机

package main

import "math"

// 动态规划可解
// 求差的方式
func maxProfit(prices []int) int {
	minPrice := math.MaxInt32
	cnt := 0
	for i := range prices {
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else if prices[i]-minPrice > cnt {
			cnt = prices[i] - minPrice
		}
	}
	return cnt
}
