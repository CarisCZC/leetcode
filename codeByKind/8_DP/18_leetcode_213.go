package main

// 现在房屋形成了环。
// 分开考虑就行
// dp[i]依然表示，i个房间内能偷盗的最多的钱
// dp[i] = max(dp[i-2]+nums[i],dp[i-1])
func rob(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	//1. 首尾都不偷的情况，但这种情况被后两种情况考虑了进去
	// a := robHelper(nums[1 : n-1])
	//2. 偷首不偷尾
	b := robHelper(nums[:n-1])
	//3. 偷尾不偷首
	c := robHelper(nums[1:])
	return max(b, c)
}
func robHelper(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	dp := make([]int, n)
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])
	for i := 2; i < n; i++ {
		dp[i] = max(dp[(i-2)]+nums[i], dp[(i-1)])
	}
	return dp[n-1]
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
