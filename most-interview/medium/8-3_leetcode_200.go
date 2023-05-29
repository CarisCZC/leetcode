// https://leetcode.cn/problems/number-of-islands/

// 200. 岛屿数量

package main

type index struct {
	x int
	y int
}

// 广度优先搜索
func numIslands(grid [][]byte) int {
	isFind := make([][]bool, len(grid))
	for i := range isFind {
		isFind[i] = make([]bool, len(grid[i]))
	}
	queue := []index{}
	res := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if !isFind[i][j] && grid[i][j] == '1' {
				queue = append(queue, index{i, j})
				for len(queue) != 0 {
					cur := queue[0]
					if !isFind[cur.x][cur.y] && grid[cur.x][cur.y] == '1' {
						isFind[cur.x][cur.y] = true
						if cur.x > 0 {
							queue = append(queue, index{cur.x - 1, cur.y})
						}
						if cur.x < len(grid)-1 {
							queue = append(queue, index{cur.x + 1, cur.y})
						}
						if cur.y > 0 {
							queue = append(queue, index{cur.x, cur.y - 1})
						}
						if cur.y < len(grid[cur.x])-1 {
							queue = append(queue, index{cur.x, cur.y + 1})
						}
					}
					queue = queue[1:]
				}
				// 队列为空时，就会有一个岛屿
				res++
			}
		}
	}
	return res
}
