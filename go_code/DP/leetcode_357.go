package main

// dp 表示有多少个每位都不相同的数
// if i 有相同的位，则dp不变，否则，dp+1
func countNumbersWithUniqueDigits(n int) int {
	if n == 0 {
		return 1
	}
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 10
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + (dp[i-1]-dp[i-2])*(10-i+1)
	}
	return dp[n]
}
