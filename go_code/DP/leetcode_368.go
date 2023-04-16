package main

import "sort"

// dp[i] 表示nums[i]为最大值能构成的最大整除子集
// dp[i] = max{dp[j],nums[i]} // nums[i]%nums[j] == 0
// 要排个序，不排序的话，往回找有点麻烦
func largestDivisibleSubset(nums []int) []int {
	sort.Ints(nums)
	dp := make([][]int, len(nums))
	// 初始化
	dp[0] = []int{nums[0]}
	i := 1
	res := dp[0]
	for ; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i]%nums[j] == 0 && (dp[i] == nil || len(dp[i]) < len(dp[j])+1) {
				tmp := make([]int, 0, len(dp[j])+1)
				tmp = append(tmp, dp[j]...)
				tmp = append(tmp, nums[i])
				dp[i] = tmp
			}
		}
		if dp[i] == nil {
			dp[i] = []int{nums[i]}
		}
		if len(dp[i]) > len(res) {
			res = dp[i]
		}
	}
	return res
}

func main() {
	largestDivisibleSubset([]int{2, 3, 4, 9, 8})
}
