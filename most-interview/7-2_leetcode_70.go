// https://leetcode.cn/problems/climbing-stairs/
// 70. 爬楼梯

package main

// n 是由 n-1，n-2组成
// 即dp[n] = dp[n-1]+dp[n-2]
// dp[1] = 1
// dp[2] = 2
func climbStairs(n int) int {
	res := 0
	dp1 := 1
	dp2 := 2
	for i := 3; i <= n; i++ {
		res = dp1 + dp2
		dp1 = dp2
		dp2 = res
	}
	if n == 2 {
		return dp2
	}
	if n == 1 {
		return dp1
	}
	return res
}
