package main

// 背包问题
// 这道题有另一种解法：
// 1. 先排序，在搜索，变成了一个回溯问题。但时间复杂度比较高
// 2. 背包解法
// 背包的容积：sum/2
// 物品的价值：nums[i]，物品的重量：nums[i]
// 要求：背包装满，或者说物品价值=sum/2（一致的）
// 每个物质只使用一次：01背包
func canPartition(nums []int) bool {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	if sum%2 != 0 {
		return false
	}
	target := sum / 2
	dp := make([]int, target+1) // 题目nums[i]<=100,len(nums)<=200
	dp[0] = 0
	for i := 0; i < len(nums); i++ {
		// 已知dp[j-1] <= dp[j]
		for j := target; j >= nums[i]; j-- {
			dp[j] = max(dp[j], dp[j-nums[i]]+nums[i])
		}
	}
	return dp[target] == target
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
