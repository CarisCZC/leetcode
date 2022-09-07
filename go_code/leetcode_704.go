// leetcode 704,二分查找
package main

func search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	mid := r / 2
	for l <= r {
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > target {
			r = mid - 1
			mid = (l + r) / 2
		} else {
			l = mid + 1
			mid = (l + r) / 2
		}
	}
	return -1
}
