// https://leetcode.cn/problems/single-number/
// 136. 只出现一次的数字
package main

// 异或运算
func singleNumber(nums []int) int {
	res := nums[0]
	for i := 1; i < len(nums); i++ {
		res ^= nums[i]
	}
	return res
}
