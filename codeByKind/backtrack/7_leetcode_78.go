// https://leetcode.cn/problems/subsets/description/

// 回溯
package main

func subsets(nums []int) [][]int {
	res := [][]int{}
	res = append(res, []int{})
	path := []int{}
	var dfs func(nums []int, start int)
	dfs = func(nums []int, start int) {
		for i := start; i < len(nums); i++ {
			path = append(path, nums[i])
			tmp := make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)
			dfs(nums, i+1)
			path = path[:len(path)-1]

		}
	}
	dfs(nums, 0)
	return res

}
