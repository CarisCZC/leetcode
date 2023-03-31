package main

// 尽可能跳的最远
// 即i+nums[i] 最大
func jump(nums []int) int {
	dp := nums[0]
	preDp := nums[0]
	res := 0
	if len(nums) == 1 {
		return res
	}
	// res++ // 进入循环，说明至少要跳一次
	for i := 1; i < len(nums); i++ {
		if i > preDp { //当目前台阶已经是前一个台阶不可达的时候，先跳到新台阶
			preDp = dp
			res++
		}
		if i <= preDp { //表示这几个台阶都是上一个跳过来的
			tmp := i + nums[i]
			if tmp > dp { // 取得最大的
				dp = tmp
			}
		}
		if preDp >= len(nums)-1 { // 此时只需要再跳一步就能到达终点了
			return res + 1
		}
	}
	return res + 1
}
