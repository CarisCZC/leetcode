package main

// 如果能组成，边长l=sum/4
// dp[i][j] 表示放入j个火柴来满足 sum = l
// 如果长度为l，则表示满足
//
func makesquare(matchsticks []int) bool {
	n := len(matchsticks)
	used := make([]bool, n) // 表示i位置的火柴已经被使用
	if n < 4 {
		return false
	}
	l := 0
	for i := range matchsticks {
		l += matchsticks[i]
	}
	if l%4 != 0 {
		return false
	}
	l /= 4
	var help func() bool
	help = func() bool {
		dp := make([]int, l+1)
		dp[0] = 0
		for i := 0; i < n; i++ {
			if !used[i] {
				for j := l; j >= matchsticks[i]; j-- {
					if dp[j-matchsticks[i]]+matchsticks[i] > dp[j] {

						dp[j] = dp[j-matchsticks[i]] + matchsticks[i]

					}
				}

			}
		}
		return dp[l] == l
	}
	for i := 0; i < 4; i++ {
		if !help() {
			return false
		}
	}
	return true
}

func main() {
	makesquare([]int{5, 5, 5, 5, 4, 4, 4, 4, 3, 3, 3, 3})
}
