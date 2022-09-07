// leetcode 892,三维形体的表面积
package main

// 每个位置都是一个柱体，计算这个位置上的柱体，所贡献的表面积
func surfaceArea(grid [][]int) int {
	n := len(grid)
	sum :=0
	for i:=0;i<n;i++{
		// colMax :=0;
		// rowMax :=0;
		for j:=0;j<len(grid[i]);j++{
			if grid[i][j] >0{//每个柱体能提供的所有面
				sum+= grid[i][j]*4+2
			}
			if i>0{
				//前一行遮挡了多少表面积
				n := min(grid[i][j], grid[i-1][j])
				sum -= n*2
			}
			if j>0{
				// 前一列遮挡了多少表面积
				m := min(grid[i][j] , grid[i][j-1])
				sum -= m*2
			}
		}
		
	}

	return sum
}

func min(a,b int)int{
	if a<b {
		return a
	}
	return b
}