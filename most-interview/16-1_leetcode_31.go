// https://leetcode.cn/problems/next-permutation/
// 31. 下一个排列

package main

// 从后向前遍历
// nums[i] <= nums[i-1],交换，此时交换是变小的
// 直到遇到nums[i] > nums[j],交换
func nextPermutation(nums []int) {
	if len(nums) <= 1 {
		return
	}
	i, j, k := len(nums)-2, len(nums)-1, len(nums)-1

	// find: A[i]<A[j]
	// 找到一个升序
	for i >= 0 && nums[i] >= nums[j] {
		i--
		j--
	}

	// 此时j~end是降序
	// 找大于A[i]的最小的A[k]
	// 交换后依然是升序
	if i >= 0 {
		for nums[i] >= nums[k] {
			k--
		}
		nums[i], nums[k] = nums[k], nums[i]
	}
	// 将j~end逆转
	for i, j := j, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}
