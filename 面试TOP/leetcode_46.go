// https://leetcode.cn/problems/permutations/?favorite=2ckc81c
package main

// 递归
// 回溯算法
func permute(nums []int) [][]int {
	res := [][]int{}
	num := make([]int, len(nums), len(nums))
	// num传进去是地址引用，任何改变都会使其他都变化
	res = create(res, num, nums, 0)
	return res
}

// 填充的数组，可用的是哪些数,当前是第几位,
func create(res [][]int, nums, candidate []int, i int) [][]int {
	//填充完毕
	if i == len(nums) {
		res = append(res, nums)
		return res
	}
	// 必须对nums进行新的拷贝，才能保证源nums不变
	nextNum := make([]int, 0, len(nums))
	nextNum = append(nextNum, nums...)
	for j := 0; j < len(candidate); j++ {
		nextNum[i] = candidate[j]
		// 必须对candidate进行新的拷贝，才能保证源candidate不变
		tmp := make([]int, 0, len(nums))
		tmp = append(tmp, candidate[:j]...)
		tmp = append(tmp, candidate[j+1:]...)
		res = create(res, nextNum, tmp, i+1)
	}
	return res
}
