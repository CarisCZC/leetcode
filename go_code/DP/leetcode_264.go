package main

// dp[i] 表示第i个丑数
// dp[i] = min(x2*2,x3*3,x5*5)
// 表示下一个丑数是上一个*质因数中的一个得到的
// p2, p3, p5初始为0，且如果dp[i] = p2*2;则p2++
func nthUglyNumber(n int) int {
	dp := make([]int, n)
	dp[0] = 1
	p2, p3, p5 := 0, 0, 0
	for i := 1; i < n; i++ {
		x2, x3, x5 := dp[p2]*2, dp[p3]*3, dp[p5]*5
		dp[i] = min(x2, min(x3, x5))
		if dp[i] == x2 {
			p2++
		}
		if dp[i] == x3 {
			p3++
		}
		if dp[i] == x5 {
			p5++
		}
	}
	return dp[n-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
