// https://leetcode.cn/problems/subsets-ii/
package main

import "sort"

// 与78基本一致，主要是去重。去重需要排序
func subsetsWithDup(nums []int) [][]int {
	path := []int{}
	res := [][]int{}
	sort.Ints(nums)
	var subset func(nums []int, start int)
	subset = func(nums []int, start int) {
		tmp := make([]int, len(path))
		copy(tmp, path)
		res = append(res, tmp)
		for i := start; i < len(nums); i++ {
			if i != start && nums[i-1] == nums[i] { //在同层循环中，不用重复数字，而下一层是否重复并不影响。
				continue
			}
			path = append(path, nums[i])
			subset(nums, i+1)
			path = path[:len(path)-1]
		}
	}
	subset(nums, 0)
	return res
}

func main() {
	subsetsWithDup([]int{1, 2, 2})
}
