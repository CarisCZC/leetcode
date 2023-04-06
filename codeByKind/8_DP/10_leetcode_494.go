package main

// dp[j] 表示装满j容积的背包，有多少种方法
// dp[j] += dp[j-nums[i]]
// 加法的总和为x，减的总和为sum-x
// x - (sum-x) = target
// x = (sum+target)/2
func findTargetSumWays(nums []int, target int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	if (sum+target)%2 != 0 {
		return 0
	}
	if target > sum || -target > sum {
		return 0
	}
	bagSize := (sum + target) / 2
	dp := make([]int, bagSize+1)
	dp[0] = 1
	for i := 0; i < len(nums); i++ {
		for j := bagSize; j >= nums[i]; j-- {
			dp[j] += dp[j-nums[i]]
		}
	}
	return dp[bagSize]
}
