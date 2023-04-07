package main

// DP[i] 表示凑到金额i的方法数
// DP[i] = DP[i-weight[i]]
// DP[0] = 1 也就是一个都不取，有一种方法
func change(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1
	for i := 0; i < len(coins); i++ {
		for j := coins[i]; j <= amount; j++ {
			dp[j] += dp[j-coins[i]]
		}
	}
	return dp[amount]
}

func main() {
	change(5, []int{1, 2, 5})
}
