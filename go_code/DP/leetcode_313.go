package main

import "math"

// 和上一个(264)思想基本一致
// dp[i] 表示第i+1个丑数
func nthSuperUglyNumber(n int, primes []int) int {
	dp := make([]int, n)
	dp[0] = 1
	numsPrime := make([]int, len(primes))
	for i := 1; i < n; i++ {
		cur, idx := math.MaxInt32, -1
		for k, v := range numsPrime {
			tmp := dp[v] * primes[k]
			if cur > tmp {
				cur = tmp
				idx = k
			} else if cur == tmp {
				// 有可能存在丑数重复的情况，比如例子中的2*7 和7*2
				// 记录第一个即可，后续的在遇到就增加指针
				numsPrime[k]++
			}
		}
		dp[i] = cur
		numsPrime[idx]++
	}
	return dp[n-1]
}

func main() {
	nthSuperUglyNumber(12, []int{2, 7, 13, 19})
}
