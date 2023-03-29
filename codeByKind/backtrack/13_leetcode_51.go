package main

// 每个皇后肯定放在单独的一排
func solveNQueens(n int) [][]string {
	path := []string{}
	res := [][]string{}
	used := make([][]bool, n, n)
	for i := range used {
		used[i] = make([]bool, n, n)
	}
	var canUse func(i, j int) bool //判断该位置是否可用
	canUse = func(i, j int) bool {
		// 行不用判断
		// 判断列
		for k := 0; k < n; k++ {
			if used[k][j] {
				return false
			}
		}
		// 根据我们的逻辑，只需判断45°和135°的对角线，因为下面的行还有放置，完全不需要判断
		// 判断左右对角线
		// 左对角线
		k1, m1 := 0, 0
		if i >= j {
			k1 = i - j
		} else { //
			m1 = j - i
		}
		for k1 < n && m1 < n {
			if used[k1][m1] {
				return false
			}
			k1++
			m1++
		}

		// 判断右对角线
		if i+j <= n-1 {
			k1 = i + j
			m1 = 0
		} else {
			k1 = n - 1
			m1 = i + j - n + 1
		}
		for k1 > -1 && m1 < n {
			if used[k1][m1] {
				return false
			}
			k1--
			m1++
		}
		return true
	}
	var dfs func(start int) // start 表示现在到了哪一行
	dfs = func(start int) {
		if len(path) == n {
			tmp := make([]string, len(path))
			copy(tmp, path)
			res = append(res, tmp)
			return
		}
		tmpRow := make([]byte, n, n)
		for i := 0; i < n; i++ {
			tmpRow[i] = '.'
		}
		for i := 0; i < n; i++ { // 这里的i表示在start这行的第几个位置
			if !canUse(start, i) {
				continue
			}
			used[start][i] = true
			tmpRow[i] = 'Q'
			path = append(path, string(tmpRow))
			dfs(start + 1)
			//回退
			used[start][i] = false
			tmpRow[i] = '.'
			path = path[:len(path)-1]
		}
	}
	dfs(0)
	return res
}

func main() {
	solveNQueens(4)
}
