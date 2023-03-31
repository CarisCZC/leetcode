// https://leetcode.cn/problems/maximum-subarray/description/
package main

import "math"

// 贪心算法
func maxSubArray(nums []int) int {
	res := math.MinInt32
	tmp := 0
	for i := 0; i < len(nums); i++ {
		tmp += nums[i]
		if tmp > res {
			res = tmp
		}
		if tmp <= 0 {
			tmp = 0
		}
	}
	return res
}

// 动态规划
func maxSubArray2(nums []int) int {
	dp := make([]int, len(nums), len(nums))
	dp[0] = nums[0] // dp[i]表示以i为结尾的最大连续子序列和
	res := dp[0]
	for i := 1; i < len(nums); i++ {
		if dp[i-1] <= 0 {
			dp[i] = nums[i]
		} else {
			dp[i] = dp[i-1] + nums[i]
		}
		if dp[i] > res {
			res = dp[i]
		}
	}
	return res
}
