// https://leetcode.cn/problems/binary-search/
package main

// 二分查找
// 要注意l=r时
func search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	// 向下取整
	mid := (l + r) / 2
	for l <= r {
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			r = mid - 1
			mid = (l + r) / 2
		} else {
			l = mid + 1
			mid = (l + r) / 2
		}
	}
	return -1
}
