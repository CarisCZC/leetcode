// https://leetcode.cn/problems/remove-duplicates-from-sorted-array/?envType=study-plan-v2&id=top-interview-150
// 3.删除有序数组中的重复项
package main

// 依旧是快慢指针
// 快指针-当前遍历元素 慢指针-当前插入位置
func removeDuplicates(nums []int) int {
	slow, fast := 0, 0
	for fast < len(nums) {
		for fast < len(nums) && nums[slow] == nums[fast] {
			fast++
		}
		slow++
		if fast < len(nums) && slow != fast {
			nums[slow] = nums[fast]
		}
	}
	return slow
}

func main() {
	removeDuplicates([]int{1, 1, 2})
}
