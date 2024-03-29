// https://leetcode.cn/problems/longest-common-subsequence/
// 1143. 最长公共子序列
package main

// dp[i][j] 表示 text1[:i],text2[:j]的最长公共子序列
// text1[i+1] == text2[j+1],则dp[i+1][j+1] = dp[i][j]+1
func longestCommonSubsequence(text1 string, text2 string) int {
	dp := make([][]int, len(text1)+1)
	for i := range dp {
		dp[i] = make([]int, len(text2)+1)
	}
	dp[0][0] = 0
	for i := 1; i <= len(text1); i++ {
		for j := 1; j <= len(text2); j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else { //继承此前的最大值
				if dp[i-1][j] > dp[i][j-1] {
					dp[i][j] = dp[i-1][j]
				} else {
					dp[i][j] = dp[i][j-1]
				}
			}
		}
	}

	return dp[len(text1)][len(text2)]
}
