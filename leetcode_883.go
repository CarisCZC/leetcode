// leetcode 883, 三维形体投影面积
package main

// 在x-y面上，看的是所有非0元素
// 在x-z面上，看到的是每行叠放最多的值
// 在y-z面上，看到的是每列的最大值
func projectionArea(grid [][]int) int {
	res :=0	
	n1,n2 := len(grid),len(grid[0])
	xz := make([]int,n1+1)//记录每行的最大值，最后一位为总值
	yz := make([]int,n2+1)//记录每列的最大值，最后一位为总值
	for i:=0;i<len(grid);i++{
		for j:=0;j<len(grid[i]);j++{
			tmp :=grid[i][j]
			if tmp != 0{
				res++
			}
			if tmp > xz[i]{
				xz[n1] += tmp-xz[i]
				xz[i] = tmp
			}
			if tmp >yz[j]{
				yz[n2] += tmp-yz[j]
				yz[j] = tmp
			}
		}
	}
	return res+xz[n1]+yz[n2]
}

//优秀解法:
//func projectionArea(grid [][]int) int {
// 	n := len(grid)
// 	sum := 0
// 	for i := 0; i < n; i++ {
// 		rowMax := 0
// 		colMax := 0
// 		for j := 0; j < n; j++ {
// 			if grid[i][j] > 0 {
// 				sum += 1
// 			}
// 			if grid[i][j] > rowMax {
// 				rowMax = grid[i][j]
// 			}
// 			if grid[j][i] > colMax {
// 				colMax = grid[j][i]
// 			}
// 		}
// 		sum += rowMax
// 		sum += colMax
// 	}
// 	return sum
// }