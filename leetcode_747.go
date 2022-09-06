// leetcode 747 , 至少是其他数字两倍的最大数

package main

//分为2部分
// 1. 找到最大数
// 2. 最大数是其他的2倍，保证是第二大的数的两倍

// 找到最大数的下标，第二大数的下标
func dominantIndex(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	first, second := nums[0], -1
	index := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] > first {
			first, second = nums[i], first
			index = i
		} else if nums[i] > second {
			second = nums[i]
		}
	}
	if first < second*2 {
		return -1
	}
	return index
}
