package main

func wiggleMaxLength(nums []int) int {
	n := len(nums)
	if n < 2 {
		return n
	}
	ans := 1
	prevDiff := nums[1] - nums[0]
	if prevDiff != 0 {
		ans++
	}
	for i := 2; i < n; i++ {
		diff := nums[i] - nums[i-1]
		if diff > 0 && prevDiff <= 0 || diff < 0 && prevDiff >= 0 {
			ans++
			prevDiff = diff
		}
	}
	return ans
}

// 动态规划解法
func wiggleMaxLength2(nums []int) int {
	n := len(nums)
	if n < 2 {
		return n
	}
	dp := make([][2]int, n)
	dp[0][0] = 1 //dp[i][0] 表示i作为波峰时的最大长度，即nums[i]>nums[i-1]
	dp[0][1] = 1 //dp[i][1] 表示i作为波谷时的最大长度，即nums[i]<nums[i-1]
	// 正常的动态规划，是到达i值时，对i之前的所有数值都进行调整
	// 而在这道题中，dp[i] >= dp[i-1],一次，只需要记录调整上一次值即可。dp[i] = dp[i-1](+1,前提时满足要求)
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[j] > nums[i] { //j>i，即num[j:i]下降
				dp[i][1] = max(dp[i][1], dp[j][0]+1)
			}
			if nums[j] < nums[i] { // j>i，nums[j:i]上升,num[j]波谷，nums[i]波峰
				dp[i][0] = max(dp[i][0], dp[j][1]+1)
			}
			if nums[j] == nums[i] { //相等
				dp[i][0] = max(dp[i][0], dp[j][0])
				dp[i][1] = max(dp[i][1], dp[j][1])
			}
		}
	}
	return max(dp[n-1][0], dp[n-1][1])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
