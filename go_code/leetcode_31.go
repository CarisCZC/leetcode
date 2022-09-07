// leetcode 31,下一个排列

package main

// 首先可以确定的是，循环应该是从后往前循环。

// 先找到一个非升序，此时i-1是较小的该换的数

// 在n-1 ~ i-1 之间找一个大于该数最小的数，然后交换

//逆转i ~ n-1

func nextPermutation(nums []int) {
	n := len(nums)
	i := n - 2
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}
	if i >= 0 {
		j := n - 1
		for j >= 0 && nums[j] <= nums[i] {
			j--
		}
		nums[j], nums[i] = nums[i], nums[j]
	}
	reverse(nums[i+1:])
}

func reverse(nums []int) {
	for i, n := 0, len(nums); i < (n / 2); i++ {
		nums[i], nums[n-i-1] = nums[n-i-1], nums[i]
	}
}
