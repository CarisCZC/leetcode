// https://leetcode.cn/problems/combination-sum/
package main

func combinationSum(candidates []int, target int) [][]int {
	res := [][]int{}
	path := []int{}
	tmpSum := 0
	var backtrack func(candidates []int, tmpSum int, target int, start int)
	backtrack = func(candidates []int, tmpSum, target int, start int) {
		for i := start; i < len(candidates); i++ {
			path = append(path, candidates[i])
			tmpSum += candidates[i]
			if tmpSum < target {
				backtrack(candidates, tmpSum, target, i)
			}
			if tmpSum == target {
				tmp := make([]int, len(path))
				copy(tmp, path)
				res = append(res, tmp)
				return
			}
			path = path[:len(path)-1]
			tmpSum -= candidates[i]
		}
	}
	backtrack(candidates, tmpSum, target, 0)
	return res
}
