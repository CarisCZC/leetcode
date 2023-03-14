package main

// 双指针经典题目
// https://leetcode.cn/problems/squares-of-a-sorted-array/
// 代码随想录 https://github.com/youngyangyang04/leetcode-master/blob/master/problems/0977.%E6%9C%89%E5%BA%8F%E6%95%B0%E7%BB%84%E7%9A%84%E5%B9%B3%E6%96%B9.md

// 1. 复数的平方会得到正数，最大值一定是在数组的两边
func sortedSquares(nums []int) []int {
	res := make([]int, len(nums), len(nums))
	l, r := 0, len(nums)-1
	for k := len(res) - 1; k > -1; k-- {
		a, b := nums[l]*nums[l], nums[r]*nums[r]
		if a >= b {
			res[k] = a
			l++
		} else {
			res[k] = b
			r--
		}
	}
	return res
}

// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }
