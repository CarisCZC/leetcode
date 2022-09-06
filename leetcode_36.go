// leetcode 36, 有效的数独

package main

func isValidSudoku(board [][]byte) bool {
	for _, row := range board {
		mp := make(map[byte]bool) //检查每一行时都初始化一下
		for _, j := range row {   //取得行
			if j != '.' {
				if !mp[j] {
					mp[j] = true
				} else {
					return false
				}
			}
		}
	}
	//检查每一列
	for i := 0; i < len(board[1]); i++ {
		mp := make(map[byte]bool)
		for j := 0; j < len(board); j++ {
			v := board[j][i]
			if v != '.' {
				if !mp[v] {
					mp[v] = true
				} else {
					return false
				}
			}
		}
	}
	//检查每个9宫格
	for i := 0; i < len(board); i += 3 { //行基数
		mp := make(map[byte]bool)
		for j := 0; j < len(board); j += 3 { //列基数
			for k := 0; k < 9; k++ {
				row := i + k/3
				col := j + k/3
				v := board[row][col%3]
				if v != '.' {
					if !mp[v] {
						mp[v] = true
					} else {
						return false
					}
				}
			}
		}

	}

	return true
}

// 官方写法：一次遍历：在每个位置时，创建了三个数组
func isValidSudoku2(board [][]byte) bool {
	var rows, cols [9][9]int  //rows[i][index]表示在第i+1行中index+1出现的次数
	var subBoxes [3][3][9]int // 将原数独分为9个9宫格，subBoxes[i][j][index]表示在下标为[i][j]的9宫格里index+1出现的次数
	for i, row := range board {
		for j, v := range row {
			if v != '.' {
				index := v - '1'
				rows[i][index]++
				cols[j][index]++
				subBoxes[i/3][j/3][index]++
				if rows[i][index] > 1 || cols[j][index] > 1 || subBoxes[i/3][j/3][index] > 1 {
					return false
				}
			}
		}
	}
	return true
}
