package main

// DP
// DP[n] 表示有多少方法到台阶n
// DP[n] = DP[n-1]+DP[n-2]
// 即有多少方法到n-1台阶以及n-2台阶
// DP[1]=1,DP[2]=2
func climbStairs(n int) int {
	if n < 3 {
		return n
	}
	dp1, dp2 := 1, 2
	for i := 3; i <= n; i++ {
		tmp := dp1 + dp2
		dp1 = dp2
		dp2 = tmp
	}
	return dp2
}
