// leetcode 628, 三个数的最大乘积

package main

import "math"

// 可以理解为，找三个最大的数
// 之前一题（414），找第三大的数，其中的想法完全可以用到这里

func maximumProduct(nums []int) int { // 处理一下负数的情况,如果有负数，必然是两负一正的情况
	if len(nums) == 3 {
		return nums[0] * nums[1] * nums[2]
	}
	first, second, third := math.MinInt32, math.MinInt32, math.MinInt32
	fuFirst, fuSecond := math.MaxInt32, math.MaxInt32

	for _, v := range nums {
		if v <= fuFirst {
			fuFirst, fuSecond = v, fuFirst
		} else if v < fuSecond {
			fuSecond = v
		}
		if v >= first {
			first, second, third = v, first, second
		} else if v >= second {
			second, third = v, second
		} else if v > third {
			third = v
		}
	}
	return max(first*second*third, fuFirst*fuSecond*first)
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
