package main

// dp[i][j] t[:j]在s[:i]中出现的个数
// dp[i][j] = dp[i-1][j]+dp[i-1][j-1] s[i-1] == t[j-1]
// dp[i][j] = dp[i-1][j] s[i-1] != t[j-1]
// dp[i] = dp[i-1]*t[i]出现的次数
func numDistinct(s string, t string) int {
	dp := make([][]int, len(s)+1)
	for i := range dp {
		dp[i] = make([]int, len(t)+1)
	}
	for i := range dp {
		dp[i][0] = 1
	}
	i := 1

	for ; i <= len(s); i++ {

		for j := 1; j <= len(t); j++ {
			if s[i-1] == t[j-1] {
				dp[i][j] = dp[i-1][j] + dp[i-1][j-1]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[i-1][len(t)]
}

func main() {
	numDistinct("rabbbit", "rabbit")
}
