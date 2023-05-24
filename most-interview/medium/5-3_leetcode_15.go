// https://leetcode.cn/problems/3sum/
// 三数之和
package main

import "sort"

// 排序+双指针
// 固定第一个元素，然后双指针找剩下两个
// 排序-双指针-去重
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	res := [][]int{}
	for i := 0; i < len(nums); i++ {
		if i != 0 && (nums[i] > 0 || nums[i] == nums[i-1]) {
			continue
		}
		cur := nums[i]
		l, r := i+1, len(nums)-1
		for l < r {
			if cur+nums[l]+nums[r] > 0 {
				r--
			} else if cur+nums[l]+nums[r] < 0 {
				l++
			} else {
				res = append(res, []int{cur, nums[l], nums[r]})
				l++
				r--
				for l < r && nums[l] == nums[l-1] {
					l++
				}
				for l < r && nums[r] == nums[r+1] {
					r--
				}
			}
		}
	}
	return res
}
