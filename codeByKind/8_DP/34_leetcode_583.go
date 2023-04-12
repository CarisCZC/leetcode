package main

// 先求公共子序列长度k，然后len(word1)-k+len(word2)-k
// dp[i][j] 表示word1[:i]与word2[:j]的公共子序列长度
// dp[i][j] = dp[i-1][j-1]+1  word1[i-1] == word2[j-1]
// dp[i][j] = max(dp[i-1][j] , dp[i][j-1])
func minDistance(word1 string, word2 string) int {
	dp := make([][]int, len(word1)+1)
	for i := range dp {
		dp[i] = make([]int, len(word2)+1)
	}
	dp[0][0] = 0
	for i := 1; i <= len(word1); i++ {
		for j := 1; j <= len(word2); j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				if dp[i][j-1] > dp[i-1][j] {
					dp[i][j] = dp[i][j-1]
				} else {
					dp[i][j] = dp[i-1][j]
				}
			}
		}
	}
	return len(word1) - dp[len(word1)][len(word2)] + len(word2) - dp[len(word1)][len(word2)]

}
