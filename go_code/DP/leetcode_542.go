package main

type point struct {
	x, y int
}

// 从右下
func updateMatrix(mat [][]int) [][]int {
	m, n := len(mat), len(mat[0])
	res := make([][]int, m)
	visit := make([][]bool, m)
	for i := range res {
		res[i] = make([]int, n)
		visit[i] = make([]bool, n)
	}
	list := []point{}
	// 先找0，然后再遍历每个位置
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if mat[i][j] == 0 {
				res[i][j] = 0
				visit[i][j] = true
				list = append(list, point{i, j})
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if !visit[i][j] {
				dist := max(list[0].x-i, i-list[0].x) + max(list[0].y-j, j-list[0].y)
				for k := 1; k < len(list); k++ {
					tmp := max(list[k].x-i, i-list[k].x) + max(list[k].y-j, j-list[k].y)
					if tmp < dist {
						dist = tmp
					}
				}
				visit[i][j] = true
				res[i][j] = dist
			}
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
