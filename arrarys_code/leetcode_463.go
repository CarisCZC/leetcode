// leetcode 463, 岛屿的周长

package main

// 尝试边缘检测
// 先找到第一个格，看其周围有几个相邻,统计其相邻格
func findNeighbor(i int, j int, grid [][]int) int {
	sum := 4
	left := j - 1
	up := i - 1
	right := j + 1
	down := i + 1
	if left >= 0 { //找左边的
		if grid[i][left] == 1 {
			sum -= 1
		}
	}
	if up >= 0 {
		if grid[up][j] == 1 {
			sum -= 1
		}
	}
	if right < len(grid[i]) {
		if grid[i][right] == 1 {
			sum -= 1
		}
	}
	if down < len(grid) {
		if grid[down][j] == 1 {
			sum -= 1
		}
	}
	return sum
}

func islandPerimeter(grid [][]int) int {
	sum := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				sum += findNeighbor(i, j, grid)
			}

		}
	}
	return sum
}
