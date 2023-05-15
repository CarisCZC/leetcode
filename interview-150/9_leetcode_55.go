// https://leetcode.cn/problems/jump-game/?envType=study-plan-v2&id=top-interview-150
// 55. 跳跃游戏

package main

// DP, 或者贪心
// 在接下来能跳的几步中，找最大的数
// dp 表示当前能达到的最大台阶数
func canJump(nums []int) bool {
	last := len(nums) - 1
	if last == 0 {
		return true
	}
	dp := nums[0]
	for i := 1; i < len(nums); i++ {
		if dp >= i { // 当前台阶能到
			tmp := i + nums[i]
			if tmp > dp {
				dp = tmp
			}
		}
		if dp >= last {
			return true
		}
	}
	return false
}
