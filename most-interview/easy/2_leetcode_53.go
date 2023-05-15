// https://leetcode.cn/problems/maximum-subarray/
// 53. 最大子数组和
package main

// dp[i] 表示到i的最大子序和
func maxSubArray(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	res := dp[0]
	for i := 1; i < len(nums); i++ {
		if dp[i-1] < 0 {
			dp[i] = nums[i]
		} else {
			dp[i] = nums[i] + dp[i-1]
		}
		if dp[i] > res {
			res = dp[i]
		}
	}
	return res
}
