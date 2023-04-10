package main

// dp[i] nums[i]结尾的最长连续子序列
// dp[i] = 1 nums[i]<=nums[i-1]
// dp[i] = dp[i-1]+1 nums[i]>nums[i-1]
func findLengthOfLCIS(nums []int) int {
	dp := make([]int, len(nums))
	res := 1
	dp[0] = 1
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			dp[i] = dp[i-1] + 1
		} else {
			dp[i] = 1
		}
		if res < dp[i] {
			res = dp[i]
		}
	}
	return res
}
