package main

func wordBreak(s string, wordDict []string) bool {
	//  1.将字典存为map,快速查找
	wordMap := map[string]bool{}
	for _, word := range wordDict {
		wordMap[word] = true
	}
	n := len(s)
	dp := make([]bool, n+1)
	// dp[i]表示s[0:i]的字符串满足要求
	dp[0] = true
	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			if _, ok := wordMap[s[j:i]]; dp[j] && ok {
				dp[i] = true
				break
			}
		}
	}

	return dp[n]

}
