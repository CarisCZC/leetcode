package main

// dp[i] 表示s[0:i]能由wordDict拼接组成
// 背包大小：s的长度
// dp[i] = dp[j]&&wordMap[s[j:i]]
// dp[0] = true
func wordBreak(s string, wordDict []string) bool {
	wordMap := make(map[string]bool, len(wordDict))
	for _, v := range wordDict {
		wordMap[v] = true
	}
	n := len(s)
	dp := make([]bool, n+1)
	dp[0] = true
	// 选词与顺序有关，因此外层是背包
	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ { // 内层才是物品
			if _, ok := wordMap[s[j:i]]; dp[j] && ok {
				dp[i] = true
				break
			}
		}
	}
	return dp[n]
}
