package main

// dp[i] 0~i的最长递增子序列长度
// dp[i] = max(dp[0~i-1])+1 nums[i]>nums[j]
// dp[i] = 1

func lengthOfLIS(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	curlen := 1
	dp[0] = curlen
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				if dp[i] < dp[j]+1 {
					dp[i] = dp[j] + 1
				}
			}
		}
		if dp[i] == 0 {
			dp[i] = 1
		}
		if dp[i] > curlen {
			curlen = dp[i]
		}
	}
	return curlen
}

func main() {
	lengthOfLIS([]int{1, 3, 6, 7, 9, 4, 10, 5, 6})
}
