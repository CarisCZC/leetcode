// leetcode 18, 四数之和
package main

import "sort"

// 与三数之和类似
func fourSum(nums []int, target int) [][]int {
	res := make([][]int, 0, len(nums))
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		n1 := nums[i]
		l := i + 1
		r := len(nums) - 1
		mid := r - 1
		for l < mid && mid < r {
			n2 := nums[l]
			n3 := nums[mid]
			n4 := nums[r]
			tmpRes := n1 + n2 + n3 + n4
			if tmpRes == target {
				res = append(res, []int{n1, n2, n3, n4})
				for l < mid && nums[mid] == n3 {
					mid--
				}
			}
			if mid <= l+1 {
				l = i + 1
				for l+1 < r && nums[r] == n4 {
					r--
				}
				mid = r - 1
				continue
			} else if tmpRes > target {
				mid--
			} else if tmpRes < target {
				l++
			}
		}
	}
	return res
}
