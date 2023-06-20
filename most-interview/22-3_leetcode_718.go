// https://leetcode.cn/problems/maximum-length-of-repeated-subarray/
// 718. 最长重复子数组

package main

// dp[i][j] 表示以nums1[:i],nums2[:j]的最长子数组的长度
func findLength(nums1 []int, nums2 []int) int {
	dp := make([][]int, len(nums1)+1)
	for i := range dp {
		dp[i] = make([]int, len(nums2)+1)
	}
	res := 0
	dp[0][0] = 0
	for i := 1; i <= len(nums1); i++ {
		for j := 1; j <= len(nums2); j++ {
			if nums1[i-1] == nums2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			}
			if dp[i][j] > res {
				res = dp[i][j]
			}
		}
	}
	return res
}
