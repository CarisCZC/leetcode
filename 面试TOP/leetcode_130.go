package main

func solve(board [][]byte) {

	m, n := len(board), len(board[0])
	needSolve := map[[2]int]bool{}
	noSolve := map[[2]int]bool{}
	var find0 func(noSolve *map[[2]int]bool, i, j int)
	find0 = func(noSolve *map[[2]int]bool, i, j int) {
		tmpSolve := *noSolve
		tmpSolve[[2]int{i, j}] = true
		// 右，左，下，上
		add := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
		for k := 0; k < len(add); k++ {
			row := i + add[k][0]
			col := j + add[k][1]
			// 不越界
			if row > -1 && row < m && col > -1 && col < n {
				if board[row][col] == 'O' && !tmpSolve[[2]int{row, col}] {
					find0(noSolve, row, col)
				}
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'O' && (i == 0 || i == m-1 || j == 0 || j == n-1) {
				if !noSolve[[2]int{i, j}] {
					find0(&noSolve, i, j)
				}
			} else {
				needSolve[[2]int{i, j}] = true
			}
		}
	}
	for k := range needSolve {
		if !noSolve[k] {
			board[k[0]][k[1]] = 'X'
		}
	}
}
