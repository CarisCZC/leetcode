package main

// DP
// DP[i]表示n=i时有多少颗二叉搜索树
// DP[0]=1,DP[1] = 1,DP[2] = 2
// DP[i] = DP[0]*DP[i-1] + DP[1]*DP[i-2]+...

func numTrees(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			dp[i] += dp[j-1] * dp[i-j]
		}
	}
	return dp[n]
}
