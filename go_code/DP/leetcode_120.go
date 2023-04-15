package main

import "math"

// 三角形的最小路径和
// dp[i][j] 表示到[i][j]时的路径
// dp[i][j] = min(dp[i-1][j],dp[i-1][j-1])+trangle[i][j]
// j从0开始遍历该层
func minimumTotal(triangle [][]int) int {
	dp := make([][]int, len(triangle))
	for i := range dp {
		dp[i] = make([]int, len(triangle[i]))
	}
	// 第一层只有一个元素
	dp[0][0] = triangle[0][0]
	i := 1
	for ; i < len(triangle); i++ {
		dp[i][0] = dp[i-1][0] + triangle[i][0] //0是继承下来的
		for j := 1; j < len(triangle[i]); j++ {
			// 上一层的dp[i-1][j-1]
			last := j
			if j == len(triangle[i])-1 {
				last = j - 1
			}
			cur := min(dp[i-1][last], dp[i-1][j-1]) + triangle[i][j]
			dp[i][j] = cur
		}
	}
	res := math.MaxInt32
	for _, k := range dp[i-1] {
		if res > k {
			res = k
		}
	}
	return res
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	minimumTotal([][]int{{-1}, {2, 3}, {1, -1, -3}})
}
