// https://leetcode.cn/problems/majority-element/
// 169. 多数元素
package main

func majorityElement(nums []int) int {
	res := nums[0]
	cnt := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == res {
			cnt++
		} else {
			cnt--
		}
		if cnt < 0 {
			res = nums[i]
			cnt = 1
		}
	}
	return res
}
