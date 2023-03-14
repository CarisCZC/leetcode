package main

// 两数之和，用map

func twoSum(nums []int, target int) (res []int) {
	resMap := map[int]int{} // k:数值 v:下标
	for i, v := range nums {
		if other, ok := resMap[target-v]; ok && other != i {
			res = []int{i, other}
			break
		}
		resMap[v] = i
	}
	return res
}
