// https://leetcode.cn/problems/longest-increasing-subsequence/
// 300. 最长递增子序列
package main

// dp[i] 表示 0~i的最长递增子序列
// dp[i+1] = max(dp[j])+1 要求nums[i+1] > nums[j]
// 要点：以i为结尾的递增序列，其最大值一定是nums[i]
func lengthOfLIS(nums []int) int {
	n := len(nums)
	dp := make([]int, len(nums))
	curlen := 1
	dp[0] = curlen
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ { //找前面的最长递增子序列
			if nums[i] > nums[j] {
				if dp[i] < dp[j]+1 {
					dp[i] = dp[j] + 1
				}
			}
		}
		if dp[i] == 0 {
			dp[i] = 1
		}
		if dp[i] > curlen {
			curlen = dp[i]
		}
	}
	return curlen
}
