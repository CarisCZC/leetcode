package main

// 用每一列来标记每行是否出现过0，用col0来标记第一列是否出现过0

func setZeroes(matrix [][]int) {
	n, m := len(matrix), len(matrix[0])
	col0 := false
	for _, r := range matrix {
		if r[0] == 0 {
			col0 = true
		}
		for j := 1; j < m; j++ {
			if r[j] == 0 {
				r[0] = 0
				// 同时第一行也要有一些变化
				matrix[0][j] = 0
			}
		}
	}
	for i := n - 1; i > -1; i-- {
		for j := 1; j < m; j++ {
			// 如果这行第一个或者这列第一个是0
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
		if col0 {
			matrix[i][0] = 0
		}
	}
}
