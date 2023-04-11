package main

// dp[i] 表示以i为结尾的最大子序和
// dp[0] = nums[0]
// dp[i] = dp[i-1]+nums[i]	nums[i]>0
// dp[i] = nums[i] nums[i]<=0
func maxSubArray(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	res := dp[0]
	for i := 1; i < len(nums); i++ {
		if dp[i-1] <= 0 {
			dp[i] = nums[i]
		} else {
			dp[i] = dp[i-1] + nums[i]
		}
		if res < dp[i] {
			res = dp[i]
		}
	}
	return res
}
