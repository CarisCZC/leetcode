// 使用最小花费爬楼梯

package main

// 自己思考过使用动态规划，结果规划失败了
func minCostClimbingStairs(cost []int) int {
	n := len(cost)
	pre, cur := 0, 0
	for i := 2; i <= n; i++ {
		pre, cur = cur, min(cur+cost[i-1], pre+cost[i-2])
	}
	return cur 
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
