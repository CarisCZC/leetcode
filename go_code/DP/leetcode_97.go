package main

// DP[i][j] 表示，由s1前i个字符+s2前j个字符能否交错组成s3前i+j个元素
// DP[i][j] = DP[i-1][j]&&s3[i+j]==s1[i] || DP[i][j-1]&&s3[i+j]==s2[j]
func isInterleave(s1 string, s2 string, s3 string) bool {
	if len(s3) != len(s1)+len(s2) {
		return false
	}
	dp := make([][]bool, len(s1)+1)
	for i := range dp {
		dp[i] = make([]bool, len(s2)+1)
	}
	dp[0][0] = true
	for i := 0; i <= len(s1); i++ {
		for j := 0; j <= len(s2); j++ {
			idx := i + j - 1
			if i > 0 {
				dp[i][j] = dp[i][j] || (dp[i-1][j] && s1[i-1] == s3[idx])
			}
			if j > 0 {
				dp[i][j] = dp[i][j] || (dp[i][j-1] && s2[j-1] == s3[idx])
			}
		}
	}
	return dp[len(s1)][len(s2)]
}

func main() {
	isInterleave("a", "b", "a")
}
