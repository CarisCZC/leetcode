// https://leetcode.cn/problems/sort-colors/?favorite=2ckc81c
package main

// 0,1,2
// 一个从前面找，一个从后面找
func sortColors(nums []int) {
	n := len(nums)
	index0 := 0
	index2 := n - 1
	for i := index0; i <= index2; {
		if nums[i] == 0 {
			nums[i], nums[index0] = nums[index0], nums[i]
			index0++
			i = index0
		} else if nums[i] == 2 {
			nums[i], nums[index2] = nums[index2], nums[i]
			index2--
			i = index0
		} else {
			i++
		}
	}
}
