// https://github.com/youngyangyang04/leetcode-master/blob/master/problems/0047.%E5%85%A8%E6%8E%92%E5%88%97II.md
package main

// used数组依然可用使用，但是额外要加一个同层map判定重复数字
func permuteUnique(nums []int) [][]int {
	path := make([]int, 0, len(nums))
	res := [][]int{}
	used := make([]bool, len(nums), len(nums))
	var dfs func(nums []int, used []bool)
	dfs = func(nums []int, used []bool) {
		if len(path) == len(nums) {
			tmp := make([]int, len(nums))
			copy(tmp, path)
			res = append(res, tmp)
			return
		}
		levelUsed := make(map[int]bool, len(nums))
		for i := 0; i < len(nums); i++ {
			if levelUsed[nums[i]] || used[i] {
				continue
			}
			path = append(path, nums[i])
			levelUsed[nums[i]] = true
			used[i] = true
			dfs(nums, used)
			used[i] = false
			path = path[:len(path)-1]
		}
	}
	dfs(nums, used)
	return res
}
