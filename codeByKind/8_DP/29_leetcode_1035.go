package main

// dp[i][j] nums1[0:i],与nums2[0:j] 能连线的最长线
// dp[i][j] = dp[i-1][j-1]+1 nums[i]==nums[j]
func maxUncrossedLines(nums1 []int, nums2 []int) int {
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
			} else {
				if dp[i-1][j] > dp[i][j-1] {
					dp[i][j] = dp[i-1][j]
				} else {
					dp[i][j] = dp[i][j-1]
				}
			}
			if res < dp[i][j] {
				res = dp[i][j]
			}
		}
	}
	return res
}
func main() {
	maxUncrossedLines([]int{2, 5, 1, 2, 5}, []int{10, 5, 2, 1, 5, 2})
}
