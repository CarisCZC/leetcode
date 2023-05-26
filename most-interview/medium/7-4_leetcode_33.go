// https://leetcode.cn/problems/search-in-rotated-sorted-array/
// 33. 搜索旋转排序数组

package main

// 其实还是二分查找
// 核心关键是在于，判断是在旋转后的左边还是右边。
// 即 nums[0] <= target && target <= nums[mid]
// 或者 nums[0] > target ,此时在旋转的右边
// 有: nums[mid] < target && target <= nums[n-1]
func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	if len(nums) == 1 {
		if nums[0] == target {
			return 0
		} else {
			return -1
		}
	}
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[0] <= nums[mid] { // 说明mid左边是有序的
			if nums[0] <= target && target <= nums[mid] { // target在该半区
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else { // mid的右边有序
			if nums[mid] <= target && target <= nums[r] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}
	return -1
}
