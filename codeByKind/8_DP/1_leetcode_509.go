package main

// DP
// DP[i] 表示第i个斐波那契数
// DP[i] =DP[i-1]+DP[i-2]
// DP[1] = 1
// DP[0] = 0
// i只和i-1和i-2有关，因此只使用两个变量即可
func fib(n int) int {
	if n < 2 {
		return n
	}
	n1, n2 := 1, 0
	for i := 2; i <= n; i++ {
		tmp := n1 + n2
		n2 = n1
		n1 = tmp
	}
	return n1
}
