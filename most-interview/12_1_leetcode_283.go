// https://leetcode.cn/problems/move-zeroes/
// 283. 移动零

package main

// 双指针
func moveZeroes(nums []int) {
	cur := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[cur] = nums[i]
			cur++
		}
	}
	for i := cur; i < len(nums); i++ {
		nums[i] = 0
	}
}
