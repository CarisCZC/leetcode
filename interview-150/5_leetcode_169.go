// https://leetcode.cn/problems/majority-element/?envType=study-plan-v2&id=top-interview-150
// 5. 多数元素
package main

// 一个计数位，一个标记位
func majorityElement(nums []int) int {
	cnt := 1
	cur := nums[0]
	for i := 1; i < len(nums); i++ {
		v := nums[i]
		if v != cur {
			cnt--
		} else {
			cnt++
		}
		if cnt < 0 {
			cur = v
			cnt = 1
		}
	}
	return cur
}
