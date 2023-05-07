package main

// 整数替换
// dp[i] 表示将i替换成1需要的最小步骤
// dp[i] = dp[i/2]+1 //i为偶数
// dp[i] = min(dp[(i-1)]+1,dp[(i+1)]+1 //i为奇数
// 结果这题DP超内存.....谢
// func integerReplacement(n int) int {
// 	dp := make([]int, n+2)
// 	// 1需要0步就可以替换
// 	dp[1] = 0
// 	for i := 2; i <= n+1; i++ {
// 		if i%2 == 0 {
// 			dp[i] = dp[i/2] + 1
// 			// 更新一下他前面的数,其前面的一定是奇数
// 			if dp[i] < dp[i-1] {
// 				dp[i-1] = dp[i] + 1
// 			}
// 		} else {
// 			dp[i] = dp[i-1] + 1
// 		}
// 	}
// 	return dp[n]
// }
func integerReplacement(n int) int {
	ans := 0
	for n != 1 {
		switch {
		case n%2 == 0:
			ans++
			n /= 2
		case n%4 == 1:
			ans += 2
			n /= 2
		case n == 3:
			ans += 2
			n = 1
		default:
			ans += 2
			n = n/2 + 1
		}
	}
	return ans
}
