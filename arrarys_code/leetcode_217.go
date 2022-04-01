// leetcode 217, 存在重复元素
package main

//没想出时间O(n)，空间O(1)的解
// go中没有set，可以用map来代替set

func containsDuplicate(nums []int) bool {
	numCount := make(map[int]bool)
	for _, v := range nums {
		if numCount[v] {
			return true
		} else {
			numCount[v] = true
		}
	}
	return false

}
