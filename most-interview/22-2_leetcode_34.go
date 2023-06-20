// https://leetcode.cn/problems/find-first-and-last-position-of-element-in-sorted-array/
// 34. 在排序数组中查找元素的第一个和最后一个位置

package main

// 找到比target大的第一个元素，找到<=target的第一个元素
func searchRange(nums []int, target int) []int {
	left := search(nums, target)
	if left == len(nums) || nums[left] != target {
		return []int{-1, -1}
	}
	right := search(nums, target+1) - 1
	return []int{left, right}
}

// 返回的是元素下标
func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}
