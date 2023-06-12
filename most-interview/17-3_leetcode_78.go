// https://leetcode.cn/problems/subsets/https://leetcode.cn/problems/subsets/
// 78. 子集

package main

func subsets(nums []int) [][]int {
	res := [][]int{}
	return helper(nums, res, []int{})
}
func helper(nums []int, res [][]int, cur []int) [][]int {
	tmp := []int{}
	tmp = append(tmp, cur...)
	res = append(res, tmp)
	if len(nums) == 0 {
		return res
	}
	for i := 0; i < len(nums); i++ {
		cur = append(cur, nums[i])
		res = helper(nums[i+1:], res, cur)
		cur = cur[:len(cur)-1]
	}
	return res
}
