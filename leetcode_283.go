// leetcode 283, 移动0

package main

// 当碰到一个不为0的数，就把他前移
func moveZeroes(nums []int) {
	numLen := 0
	for _, v := range nums {
		if v != 0 {
			nums[numLen] = v
			numLen++
		}
	}
	for ; numLen < len(nums); numLen++ {
		nums[numLen] = 0
	}
}
