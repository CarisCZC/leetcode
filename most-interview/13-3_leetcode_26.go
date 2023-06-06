// https://leetcode.cn/problems/remove-duplicates-from-sorted-array/description/
// 26. 删除有序数组中的重复项
package main

// 快慢指针
func removeDuplicates(nums []int) int {
	cur := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[cur-1] {
			nums[cur] = nums[i]
			cur++
		}
	}
	return cur
}
