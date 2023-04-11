package main

// 这题其实应该在29题前，两者思路基本一致
// dp[i][j] 表示text1[:i]和text2[:j]中公共子序列的最大长度
// text1[i-1] == text2[i-1]时，dp[i][j] = dp[i-1][j-1]+1
func longestCommonSubsequence(text1 string, text2 string) int {
	dp := make([][]int, len(text1)+1)
	for i := range dp {
		dp[i] = make([]int, len(text2)+1)
	}
	// 初始化
	dp[0][0] = 0
	for i := 1; i <= len(text1); i++ {
		for j := 1; j <= len(text2); j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				// 在不相等的情况下，应该继承此前的最大值
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

func main() {
	longestCommonSubsequence("abcde", "ace")
}
