// https://leetcode.cn/problems/remove-duplicates-from-sorted-array-ii/?envType=study-plan-v2&id=top-interview-150
// 4. 删除有序数组中的重复项 II
package main

// 快慢指针还是适用，只不过要多一个标记
func removeDuplicates(nums []int) int {
	n := len(nums)
	if n <= 2 {
		return n
	}
	slow, fast := 2, 2
	for fast < n {
		if nums[slow-2] != nums[fast] { // slow是待插入的位置，fast是待插入的元素。
			// slow前面的一定<=slow。因此只需要判断slow前两位是否等于fast即可
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	return slow
}
