// https://leetcode.cn/problems/combination-sum/
// 39. 组合总和
package main

func combinationSum(candidates []int, target int) [][]int {
	res := helper(candidates, []int{}, 0, target, [][]int{})
	return res
}

func helper(candidates []int, path []int, cur, target int, res [][]int) [][]int {
	if cur == target {
		tmp := []int{}
		tmp = append(tmp, path...)
		res = append(res, tmp)
		return res
	}
	for i := 0; i < len(candidates); i++ {
		tmp := cur + candidates[i]
		if tmp > target {
			continue
		}
		path = append(path, candidates[i])
		res = helper(candidates[i:], path, tmp, target, res)
		path = path[:len(path)-1]
	}
	return res
}

func main() {
	combinationSum([]int{2, 3, 6, 7}, 7)
}
