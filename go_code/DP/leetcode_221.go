package main

// dp[i][j] 表示[i,j]为右下角的全为1的正方形的最大边长
// 为什么是右下角而不是左上角？
// dp的原则是后一个状态要由前一状态推算出来，如果是左上角，相当于前一状态由后面的状态推出
func maximalSquare(matrix [][]byte) int {
	m, n := len(matrix), len(matrix[0])
	dp := make([][]int, m)
	res := 0
	for i := range dp {
		dp[i] = make([]int, n)
		dp[i][0] = int(matrix[i][0] - '0')
		if dp[i][0] > res {
			res = dp[i][0]
		}
	}
	//要初始化第一行第一列，第一行第一列最大值取决于自身值
	for i := 0; i < n; i++ {
		dp[0][i] = int(matrix[0][i] - '0')
		if dp[0][i] > res {
			res = dp[0][i]
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][j] == '0' {
				dp[i][j] = int(matrix[i][j] - '0')
			} else {
				// 最小值可能为0
				dp[i][j] = min(dp[i-1][j-1], min(dp[i-1][j], dp[i][j-1])) + 1
			}
			if dp[i][j] > res {
				res = dp[i][j]
			}
		}
	}
	return res * res
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	maximalSquare([][]byte{
		{'1', '0', '1', '0', '0'},
		{'1', '0', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '0', '0', '1', '0'},
	},
	)
}
