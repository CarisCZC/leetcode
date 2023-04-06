package main

// 背包大小:m+n
// 选择要求：sum(子集中1和0的总和) <= m+n 且sum(0) <= m
func findMaxForm(strs []string, m int, n int) int {
	// dp[i][j] 表示，有i个0和j个1的最大子集的大小。
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for _, v := range strs {
		oneNum, zeroNum := 0, 0
		for _, char := range v {
			if char == '0' {
				zeroNum++
			} else {
				oneNum++
			}
		}
		for i := m; i >= zeroNum; i-- {
			for j := n; j >= oneNum; j-- {
				dp[i][j] = max(dp[i][j], dp[i-zeroNum][j-oneNum]+1)
			}
		}
	}
	return dp[m][n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
