// leetcode 27, 移除元素
package main

func removeElement(nums []int, val int) int {
	tmpLen := 0
	for _, v := range nums {
		if v != val {
			nums[tmpLen] = v
			tmpLen++
		}

	}
	return tmpLen

}
