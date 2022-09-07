// leetcode 643, 子数组的最大平均数Ⅰ
package main

// 计算0~（k-1）的和sum以及平均数ans
// sum-nums[0]+num[k],并计算tmpans，如果大于ans，则替换
func findMaxAverage(nums []int, k int) float64 {
	sum := 0
	for i := 0; i < k; i++ {
		sum += nums[i]
	}
	ans := sum
	for i := k; i < len(nums); i++ {
		sum = sum - nums[i-k] + nums[i]
		if sum > ans {
			ans = sum
		}
	}
	return float64(sum) / float64(k)
}
