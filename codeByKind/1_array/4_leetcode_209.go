package main

import "math"

// 滑动窗口
// 所谓滑动窗口，就是不断的调节子序列的起始位置和终止位置，从而得出我们要想的结果。
// 滑动窗口也是双指针的一种
// 要点： 1. 先调整终止位置，以达到要求
// 2. 再调整起始位置，缩小范围
func minSubArrayLen(target int, nums []int) int {
	res := math.MaxInt32
	j := 0 // j表示长度
	tmpSum := 0
	for i := 0; i+j < len(nums); {
		tmpSum += nums[i+j]
		j++
		// 放缩i
		for tmpSum >= target {
			if j < res {
				res = j
			}
			tmpSum -= nums[i]
			i++
			j--
		}
	}
	if res == math.MaxInt32 {
		return 0
	}
	return res
}
