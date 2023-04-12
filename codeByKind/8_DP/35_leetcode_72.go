package main

// 编辑距离
// 子序列在word1里可以不连续，但在word2中要求是连续的
// word2的连续子序列在word1中的最大长度
// dp[i][j] 表示，word2[:j]在word1[:i]中的最长公共子序列
// word2[j-1] == word1[i-1] dp[i][j] = dp[i-1][j-1]+1
// word2[j-1] != word1[i-1] dp[i][j] = 1
func minDistance(word1 string, word2 string) int {
	n, m := len(word1), len(word2)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
	}
	for i := range dp {
		dp[i][0] = i
	}
	for j := 0; j <= m; j++ {
		dp[0][j] = j
	}
	dp[0][0] = 0
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1] // 不操作
			} else {
				// dp[i][j] = dp[i-1][j-1] + 1 //替换
				// dp[i][j] = dp[i-1][j] + 1   // 增加一个元素
				// dp[i][j] = dp[i][j-1] + 1   // word2删除一个元素,即word1删除一个元素
				dp[i][j] = min(dp[i-1][j-1]+1, min(dp[i-1][j]+1, dp[i][j-1]+1))
			}
		}

	}
	return dp[n][m]
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	minDistance("intention", "execuetion")
}
