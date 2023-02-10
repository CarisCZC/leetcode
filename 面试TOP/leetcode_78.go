// https://leetcode.cn/problems/subsets/?favorite=2ckc81c
package main

func subsets(nums []int) [][]int {
	res := [][]int{}
	res = append(res, []int{})
	cur := make([]int, 0, len(nums))
	for i := 1; i < len(nums); i++ {
		res = subset(res, nums, i, cur)
	}
	res = append(res, nums)
	return res
}

func subset(res [][]int, nums []int, n int, cur []int) [][]int {
	if n == 0 {
		// res =
		return append(res, cur)
	}

	for i := 0; i < len(nums); i++ {
		tmp := []int{}
		tmp = append(tmp, cur...)
		tmp = append(tmp, nums[i])
		res = subset(res, nums[i+1:], n-1, tmp)
	}
	return res
}
