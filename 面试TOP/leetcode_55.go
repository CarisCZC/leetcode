// https://leetcode.cn/problems/jump-game/?favorite=2ckc81c
package main

// 尽可能能跳远
// 官方写法，比自己写的好
func canJump(nums []int) bool {
	n := len(nums)
	rightmost := 0
	for i := 0; i < n; i++ {
		// 当前i的位置，就是最佳位置
		if i <= rightmost {
			rightmost = max(rightmost, i+nums[i])
			if rightmost >= n-1 {
				return true
			}
		}
	}
	return false
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
