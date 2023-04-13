package main

// dp表示s构成回文串的数量
func countSubstrings(s string) int {

	// dp[i][j] 表示s[i:j]构成回文串
	dp := make([][]bool, len(s))
	for i := range dp {
		dp[i] = make([]bool, len(s)+1)
		dp[i][i] = true
	}
	res := 0
	for i := len(s) - 1; i > -1; i-- {
		for j := i; j < len(s); j++ {
			if s[i] == s[j] {
				if j-i <= 1 {
					res++
					dp[i][j] = true
				} else if dp[i+1][j-1] {
					res++
					dp[i][j] = true
				}
			}
		}
	}
	return res
}
