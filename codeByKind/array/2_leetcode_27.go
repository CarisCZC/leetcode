package main

// 代码随想录 https://github.com/youngyangyang04/leetcode-master/blob/master/problems/0027.%E7%A7%BB%E9%99%A4%E5%85%83%E7%B4%A0.md
// 双指针
// 双指针法（快慢指针法）： 通过一个快指针和慢指针在一个for循环下完成两个for循环的工作。
// 快指针：寻找新数组的元素 ，新数组就是不含有目标元素的数组
// 慢指针：指向更新 新数组下标的位置
func removeElement(nums []int, val int) int {
	slow, fast := 0, 0
	for fast < len(nums) {
		if nums[slow] != val { //该值要被覆盖
			slow++
			fast++
		} else {
			for nums[fast] == val { //寻找下一个非val值
				fast++
			}
		}
		if fast < len(nums) && slow != fast {
			nums[slow] = nums[fast]
		}
	}
	return slow
}
