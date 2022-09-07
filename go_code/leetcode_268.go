// leetcode 268, 丢失的数字

// 因为一定是0~n，所以所有数加起来就是n(n+1)/2，然后减去所有已存在的数即可

package main

func missingNumber(nums []int) int {
	allNum := (len(nums) * (len(nums) + 1)) >> 1
	for _, v := range nums {
		allNum -= v
	}
	return allNum
}
