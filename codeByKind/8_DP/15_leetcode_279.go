package main

// dp[j]表示，组成j的完全平方数的最小数量
// dp[j] = min(dp[j-nums[i]]+1,j^1/2)
func numSquares(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1 //0 = 0^2，即一个数

	for i := 1; i <= n; i++ {
		tmp := i * i
		if tmp > n {
			break
		}
		for j := tmp; j <= n; j++ {
			if j == tmp {
				dp[j] = 1
				continue
			}
			if dp[j] == 0 || dp[j] > dp[j-tmp]+1 {
				dp[j] = dp[j-tmp] + 1
			}
		}
	}
	return dp[n]
}

func main() {
	numSquares(12)
}
