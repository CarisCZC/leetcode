// https://leetcode.cn/problems/remove-element/?envType=study-plan-v2&id=top-interview-150
// 2. 移除元素

package main

// 双指针： 慢指针指向当前需要覆盖的位置，快指针指向当前遍历元素
func removeElement(nums []int, val int) int {
	slow, fast := 0, 0
	for fast < len(nums) {
		if nums[slow] != val {
			slow++
		}
		fast++
		if fast < len(nums) && slow != fast {
			nums[slow] = nums[fast]
		}
	}
	return slow
}
