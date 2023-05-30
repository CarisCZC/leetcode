// https://leetcode.cn/problems/longest-palindromic-substring/
// 5. 最长回文子串

package main

// dp
// dp[i][j] i~j回文字串的长度
// dp[i-1][j+1]，如果dp[i][j]==回文字串，且s[i-1]==s[j+1]
func longestPalindrome(s string) string {
	dp := isPalindrome(s)
	start, end := -1, -1
	maxLen := 0
	for i := len(s) - 1; i > -1; i-- {
		for j := i; j < len(s); j++ {
			tmpLen := j - i + 1
			if dp[i][j] && tmpLen > maxLen {
				maxLen = tmpLen
				start = i
				end = j
			}
		}
	}
	return s[start : end+1]
}

func isPalindrome(s string) [][]bool {
	dp := make([][]bool, len(s))
	for i := range dp {
		dp[i] = make([]bool, len(s))
		dp[i][i] = true
	}
	for i := len(s) - 1; i > -1; i-- {
		for j := i + 1; j < len(s); j++ {
			if j-i == 1 {
				dp[i][j] = s[i] == s[j]
			} else {
				dp[i][j] = dp[i+1][j-1] && s[i] == s[j]
			}
		}
	}
	return dp
}

func main() {
	longestPalindrome("cbbd")
}
