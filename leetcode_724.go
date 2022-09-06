// leetcode 724, 寻找中心下标

package main

// 还是前缀和的形式更好一点
func pivotIndex(nums []int) int {
	tmp := nums[0]
	for i := 1; i < len(nums); i++ {
		nums[i] += tmp
		tmp = nums[i]
	}
	for i := 0; i < len(nums); i++ {
		left := 0
		if i > 0 {
			left = nums[i-1]
		}
		if left == nums[len(nums)-1]-nums[i] {
			return i
		}
	}
	return -1
}
