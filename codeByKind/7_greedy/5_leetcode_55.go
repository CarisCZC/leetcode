package main

// 经典DP
// dp[i] 表示在该台阶能达到的最大长度
// dp[i] = nums[i]
// max = i+num[i]
// 但前提是能到i这个台阶
// dp[i]>=n-1 说明可用到最后一个台阶
// 只和i-1有关，可用滚动变量
func canJump(nums []int) bool {
	last := len(nums) - 1
	if last == 0 {
		return true
	}
	// max := 0
	dp := nums[0]
	for i := 1; i < len(nums); i++ {
		// 能到该台阶
		if dp >= i {
			// 是否能跳更远
			tmp := i + nums[i]
			if tmp > dp { //选择更远
				dp = tmp
			}
		}
		if dp >= last {
			return true
		}
	}
	return false
}
