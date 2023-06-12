// https://leetcode.cn/problems/coin-change/
// 322. 零钱兑换
package main

// 背包问题。而且是完全背包
// 背包大小11
// dp[j] 表示凑成11的最小硬币数，初始化为-1
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = -1
	}
	dp[0] = 0
	for i := 0; i < len(coins); i++ {
		for j := coins[i]; j <= amount; j++ {
			if dp[j-coins[i]] >= 0 {
				if dp[j] < 0 || dp[j] > dp[j-coins[i]]+1 {
					dp[j] = dp[j-coins[i]] + 1
				}
			}
		}
	}
	return dp[amount]
}
