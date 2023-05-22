// https://leetcode.cn/problems/two-sum/
// 1. 两数之和
package main

// 经典map题
func twoSum(nums []int, target int) (res []int) {
	another := make(map[int]int) // key:nums[i]，value:i
	for i, v := range nums {
		if anotherIndex, ok := another[target-v]; ok {
			res = []int{i, anotherIndex}
			return
		} else {
			another[v] = i
		}
	}
	return
}
