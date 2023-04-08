package main

// dp[i] 表示i个房间内能偷的最多的钱
// 如果偷i，则i-1一定没有被偷。如果不偷i，则考虑前i-1个房间最多偷多少
// DP[i] = max(dp[i-2]+nums[i],dp[i-1])
func rob(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	dp := make([]int, n)
	dp[0] = nums[0]
	dp[1] = max(dp[0], nums[1]) // 前两个房间里的最大值
	for i := 2; i < n; i++ {
		dp[i] = max(dp[i-2]+nums[i], dp[i-1])
	}
	return dp[n-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
