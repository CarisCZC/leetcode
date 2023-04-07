package main

// DP[i][j] 从nums的前i个数中任取，组成j的组合个数
// DP[0][0] = 1
// DP[i][j] += dp[i-1][j-nums[i]]
// 1维数组即可
func combinationSum4(nums []int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 1
	for j := 0; j <= target; j++ {
		for i := 0; i < len(nums); i++ {
			if j-nums[i] >= 0 {
				dp[j] += dp[j-nums[i]]
			}
		}
	}
	return dp[target]
}
