package main

// 解数独
func solveSudoku(board [][]byte) {
	var isValid func(row, col int, char byte) bool
	isValid = func(row, col int, char byte) bool {
		for i := 0; i < 9; i++ {
			if board[row][i] == char || board[i][col] == char { //同行重复
				return false
			}
		}
		// 3*3 重复
		startRow, startCol := (row/3)*3, (col/3)*3
		for i := startRow; i < startRow+3; i++ {
			for j := startCol; j < startCol+3; j++ {
				if board[i][j] == char {
					return false
				}
			}
		}
		return true
	}
	var backtarck func() bool
	backtarck = func() bool {
		for i := 0; i < 9; i++ {
			for j := 0; j < 9; j++ {
				if board[i][j] != '.' {
					continue
				}
				for char := '1'; char <= '9'; char++ {
					if isValid(i, j, byte(char)) {
						board[i][j] = byte(char)
						if backtarck() {
							return true
						}
						board[i][j] = '.'
					}
				}
				return false // 9个数字尝试了，但没有得到合适的结果，说明数字都不行，返回false
			}
		}
		return true // 没有返回false，说明得到了正确解
	}
	backtarck()
}

func main() {
	solveSudoku([][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	})
}
