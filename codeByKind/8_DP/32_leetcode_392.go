package main

// dp[i]表示s[:i]是t的子序列
// dp[i] = dp[i-1]&s[j]==t[i],
func isSubsequence(s string, t string) bool {
	dp := make([]bool, len(s)+1)
	j := 1
	dp[0] = true //空集是t的子序列
	for i := 0; i < len(t); i++ {
		if j >= len(s) {
			break
		}
		dp[j] = dp[j-1] && s[j-1] == t[i]
		if dp[j] {
			j++
		}
	}
	return dp[len(dp)-1]
}
