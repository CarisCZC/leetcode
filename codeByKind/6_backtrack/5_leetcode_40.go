package main

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	// 先排序
	sort.Ints(candidates)
	res := [][]int{}
	path := []int{}
	tmpSum := 0
	var backtrack func(candidates []int, target int, tmpSum int)
	backtrack = func(candidates []int, target, tmpSum int) {
		for i := 0; i < len(candidates); i++ {
			if i != 0 && candidates[i] == candidates[i-1] { // 去重
				continue
			}
			if candidates[i] > target-tmpSum || tmpSum > target { //剪枝
				return
			}
			path = append(path, candidates[i])
			tmpSum += candidates[i]
			if tmpSum == target {
				tmp := make([]int, len(path))
				copy(tmp, path)
				res = append(res, tmp)
			} else {
				backtrack(candidates[i+1:], target, tmpSum)
			}
			// 回退
			tmpSum -= candidates[i]
			path = path[:len(path)-1]
		}
	}
	backtrack(candidates, target, tmpSum)
	return res
}

func main() {
	combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8)
}
