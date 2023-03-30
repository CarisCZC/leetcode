// https://leetcode.cn/problems/permutations/?favorite=2ckc81c
package main

// 也可以使用一个used数组来标记，当前还有哪些值没有使用
func permute(nums []int) [][]int {
	path := []int{}
	res := [][]int{}
	used := make([]bool, len(nums), len(nums))
	var backtrack func(nums []int, used []bool)
	backtrack = func(nums []int, used []bool) {
		if len(path) == len(nums) {
			tmp := make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)
			return
		}
		for i := 0; i < len(nums); i++ {
			if used[i] {
				continue
			}
			path = append(path, nums[i])
			used[i] = true
			backtrack(nums, used)
			used[i] = false
			path = path[:len(path)-1]
		}
	}
	backtrack(nums, used)
	return res

}
