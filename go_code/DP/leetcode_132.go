package main

// 分割回文串2
// https://leetcode.cn/problems/palindrome-partitioning-ii/description/

// 从中心延展
// f[i][j] 表示从i到j是一个回文字符串
// dp[i]表示以0开始，s[i]为结尾的的字符串，最少分割次数为dp[i]
func minCut(s string) int {
	n := len(s)
	f := make([][]bool, n)
	for i := range f {
		f[i] = make([]bool, n)
		for j := range f[i] {
			f[i][j] = true
		}
	}
	for i := n - 1; i > -1; i-- {
		for j := i + 1; j < n; j++ {
			f[i][j] = s[i] == s[j] && f[i+1][j-1]
		}
	}

	dp := make([]int, n)
	for i := range dp {
		dp[i] = i
	}
	for i := range s {
		// 即0~i不需要分割
		if f[0][i] {
			dp[i] = 0
			continue
		}
		for j := 0; j < i; j++ {
			if f[j+1][i] {
				if dp[i] > dp[j]+1 {
					dp[i] = dp[j] + 1
				}
			}
		}
	}
	return dp[n-1]
}
