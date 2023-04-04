package main

//这题可以贪心，每次都选最小代价的，最后代价一定最小
// DP[n] 表示到台阶n的最小代价
// DP[n] = min(dp[n-1],dp[n-2])+min(cost[n-1],cost[n-2])
// 即前面的最小代价加上当前台阶的代价
// 初始化：DP[0] = cost[0],DP[1] = cost[1]
func minCostClimbingStairs(cost []int) int {
	dp := make([]int, len(cost))
	dp[0] = 0
	dp[1] = 0
	for i := 2; i < len(cost); i++ {
		dp[i] = min(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])
	}
	n := len(cost)
	return min(dp[n-1]+cost[n-1], dp[n-2]+cost[n-2])
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	minCostClimbingStairs([]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1})
}
