package main

func longestPalindromeSubseq(s string) int {
	// dp[i][j] 表示s[i:j]的最长回文子序列
	dp := make([][]int, len(s))
	for i := range dp {
		dp[i] = make([]int, len(s))
		dp[i][i] = 1
	}
	for i := len(s) - 1; i > -1; i-- {
		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				if dp[i][j-1] > dp[i+1][j] {
					dp[i][j] = dp[i][j-1]
				} else {
					dp[i][j] = dp[i+1][j]
				}
			}
		}
	}
	return dp[0][len(s)-1]
}
