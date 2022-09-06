// leetcode 561, 数组拆分
package main

import "sort"

// 将数组排序，然后返回偶数位的数的和
func arrayPairSum(nums []int) int {
	sort.Ints(nums)
	ans := 0
	for i := 0; i < len(nums); i += 2 {
		ans += nums[i]
	}
	return ans

}
