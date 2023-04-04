package main

// 背包问题
// weight = nums[i];value = nums[i]
// 背包容积 target=sum/2
// 每次选择一定有一个石头消失，假设消失的重量为neg
// 则剩余重量= (sum-neg)-neg，即sum-2neg。
// 那么neg就在小于sum/2的前提下，尽可能的大
func lastStoneWeightII(stones []int) int {
	sum := 0
	for _, v := range stones {
		sum += v
	}
	target := sum / 2
	dp := make([]int, target+1)

	for i := 0; i < len(stones); i++ {
		for j := target; j >= stones[i]; j-- {
			if dp[j] < dp[j-stones[i]]+stones[i] {
				dp[j] = dp[j-stones[i]] + stones[i]
			}
		}
	}
	return sum - 2*dp[target]
}
