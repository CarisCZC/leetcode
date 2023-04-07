package main

// dp[j],表示凑成金额j的最小硬币数
// dp[j] = min(dp[j],dp[j-nums[i])
// dp[0] = 1
// dp得初始化为-1
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

func main() {
	coinChange([]int{2, 5, 10, 1}, 27)
}
