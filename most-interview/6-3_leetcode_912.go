package main

import (
	"math/rand"
)

// 912. 排序数组
// 手撕快排
func sortArray(nums []int) []int {
	partition(nums, 0, len(nums)-1)
	return nums
}

func partition(nums []int, start, end int) {
	if start >= end {
		return
	}
	randomIdx := rand.Intn(end-start+1) + start
	nums[start], nums[randomIdx] = nums[randomIdx], nums[start]
	l, r := start, end
	for l < r { // 哨兵固定为nums[l]，但nums[l]具体是谁并不固定
		for l < r && nums[r] > nums[l] {
			r--
		}
		if l < r {
			nums[l], nums[r] = nums[r], nums[l]
			l++
		}
		for l < r && nums[l] < nums[r] {
			l++
		}
		// 此时nums[l] >pos 或者l=r
		if l < r {
			nums[r], nums[l] = nums[l], nums[r]
			r--
		}
	}
	partition(nums, start, l-1)
	partition(nums, r+1, end)
	return

}
